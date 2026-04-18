package routeros

import (
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
type BulkReadCache struct {
	mu    sync.RWMutex
	items map[string]map[string]MikrotikItem

	// group coalesces concurrent bulk fetches for the same path so that N
	// parallel Read calls trigger exactly one network round-trip.
	group singleflight.Group
}

func NewBulkReadCache() *BulkReadCache {
	return &BulkReadCache{items: map[string]map[string]MikrotikItem{}}
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

// populate replaces the cache entry for path with the provided items, indexed
// by their RouterOS .id. Items without a .id are skipped defensively.
func (c *BulkReadCache) populate(path string, items []MikrotikItem) {
	bucket := make(map[string]MikrotikItem, len(items))
	for _, it := range items {
		if id := it.GetID(Id); id != "" {
			bucket[id] = it
		}
	}
	c.mu.Lock()
	c.items[path] = bucket
	c.mu.Unlock()
}

// Invalidate drops the cache entry for path. Safe to call on unknown paths.
func (c *BulkReadCache) Invalidate(path string) {
	c.mu.Lock()
	delete(c.items, path)
	c.mu.Unlock()
}

// DoBulkFetch runs fetch at most once per concurrent burst of callers for the
// same path (singleflight semantics). Used by ReadItems to avoid thundering-herd
// bulk GETs during parallel refresh.
func (c *BulkReadCache) DoBulkFetch(path string, fetch func() ([]MikrotikItem, error)) ([]MikrotikItem, error) {
	v, err, _ := c.group.Do(path, func() (any, error) {
		items, err := fetch()
		if err != nil {
			return nil, err
		}
		c.populate(path, items)
		return items, nil
	})
	if err != nil {
		return nil, err
	}
	return v.([]MikrotikItem), nil
}
