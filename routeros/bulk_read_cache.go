package routeros

import (
	"errors"
	"strconv"
	"sync"

	"golang.org/x/sync/singleflight"
)

// BulkReadCache stores one-shot bulk reads of a RouterOS resource path so that
// subsequent per-id reads during the same provider operation are served from
// memory instead of issuing individual filtered GETs.
//
// The cache is opt-in via the provider-level "bulk_read_refresh" attribute. It
// is populated on the first ReadItems call for a path and invalidated on any
// successful Create/Update/Delete that targets that path.
//
// Concurrency model: a per-path generation counter is bumped by Invalidate.
// DoBulkFetch captures the generation at fetch-start and refuses to populate
// if it changed mid-flight, so writes that race with a bulk fetch cannot
// leave the cache holding a pre-write snapshot. Singleflight keys embed the
// generation, so writers that bump the generation do not coalesce with an
// already-in-flight stale fetch.
type BulkReadCache struct {
	mu    sync.RWMutex
	items map[string]map[string]MikrotikItem
	gen   map[string]uint64

	group singleflight.Group
}

// errStaleBulkFetch signals that Invalidate ran while a bulk fetch was in
// flight, so the fetched items predate the concurrent write and must be
// discarded. DoBulkFetch catches this sentinel internally and retries.
var errStaleBulkFetch = errors.New("bulk read: cache invalidated during fetch")

const maxBulkFetchAttempts = 8

func NewBulkReadCache() *BulkReadCache {
	return &BulkReadCache{
		items: map[string]map[string]MikrotikItem{},
		gen:   map[string]uint64{},
	}
}

// Lookup reports whether a bulk fetch for path has been done (pathCached) and,
// if so, whether the given id is present (itemFound). A pathCached=false result
// means the caller must trigger a bulk fetch via DoBulkFetch.
func (c *BulkReadCache) Lookup(path, id string) (item MikrotikItem, itemFound bool, pathCached bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	bucket, ok := c.items[path]
	if !ok {
		return nil, false, false
	}
	item, itemFound = bucket[id]
	return item, itemFound, true
}

// populateIfFresh replaces the cache entry for path only when the generation
// observed at fetch-start still matches. Returns true on successful populate.
func (c *BulkReadCache) populateIfFresh(path string, items []MikrotikItem, startGen uint64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.gen[path] != startGen {
		return false
	}
	bucket := make(map[string]MikrotikItem, len(items))
	for _, it := range items {
		if id := it.GetID(Id); id != "" {
			bucket[id] = it
		}
	}
	c.items[path] = bucket
	return true
}

// Invalidate drops the cache entry for path and bumps its generation so any
// concurrently in-flight bulk fetch observes the change and refuses to populate
// a stale snapshot. Safe to call on unknown paths.
func (c *BulkReadCache) Invalidate(path string) {
	c.mu.Lock()
	delete(c.items, path)
	c.gen[path]++
	c.mu.Unlock()
}

// currentGen returns the path's current generation under read lock.
func (c *BulkReadCache) currentGen(path string) uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.gen[path]
}

// DoBulkFetch runs fetch at most once per concurrent burst of callers at the
// same generation. If Invalidate runs during the fetch, the snapshot is
// discarded (errStaleBulkFetch) and the caller retries at the new generation;
// writers therefore do not wait on a stale in-flight fetch.
func (c *BulkReadCache) DoBulkFetch(path string, fetch func() ([]MikrotikItem, error)) ([]MikrotikItem, error) {
	for attempt := 0; attempt < maxBulkFetchAttempts; attempt++ {
		startGen := c.currentGen(path)
		key := path + ":" + strconv.FormatUint(startGen, 10)
		v, err, _ := c.group.Do(key, func() (any, error) {
			items, err := fetch()
			if err != nil {
				return nil, err
			}
			if !c.populateIfFresh(path, items, startGen) {
				return nil, errStaleBulkFetch
			}
			return items, nil
		})
		if errors.Is(err, errStaleBulkFetch) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return v.([]MikrotikItem), nil
	}
	return nil, errors.New("bulk read: cache repeatedly invalidated during fetch, aborting")
}
