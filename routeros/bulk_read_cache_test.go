package routeros

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBulkReadCache_Lookup_EmptyCache(t *testing.T) {
	c := NewBulkReadCache()
	item, itemFound, pathCached := c.Lookup("/ip/dns/static", "*1")
	if pathCached {
		t.Errorf("pathCached = true, want false for empty cache")
	}
	if itemFound {
		t.Errorf("itemFound = true, want false for empty cache")
	}
	if item != nil {
		t.Errorf("item = %v, want nil", item)
	}
}

// primeCache populates the cache via a synchronous DoBulkFetch so tests can
// set up cache state without duplicating the production populate path.
func primeCache(t *testing.T, c *BulkReadCache, path string, items []MikrotikItem) {
	t.Helper()
	if _, err := c.DoBulkFetch(path, func() ([]MikrotikItem, error) { return items, nil }); err != nil {
		t.Fatalf("primeCache: %v", err)
	}
}

func TestBulkReadCache_PopulateAndLookup_Hit(t *testing.T) {
	c := NewBulkReadCache()
	primeCache(t, c, "/ip/dns/static", []MikrotikItem{
		{".id": "*1", "name": "alpha"},
		{".id": "*2", "name": "beta"},
	})
	item, itemFound, pathCached := c.Lookup("/ip/dns/static", "*2")
	if !pathCached {
		t.Errorf("pathCached = false, want true")
	}
	if !itemFound {
		t.Errorf("itemFound = false, want true")
	}
	if item["name"] != "beta" {
		t.Errorf("item[name] = %v, want beta", item["name"])
	}
}

func TestBulkReadCache_PopulateAndLookup_MissOnKnownPath(t *testing.T) {
	c := NewBulkReadCache()
	primeCache(t, c, "/ip/dns/static", []MikrotikItem{{".id": "*1"}})
	item, itemFound, pathCached := c.Lookup("/ip/dns/static", "*999")
	if !pathCached {
		t.Errorf("pathCached = false, want true (path was populated)")
	}
	if itemFound {
		t.Errorf("itemFound = true, want false (id not in bulk response)")
	}
	if item != nil {
		t.Errorf("item = %v, want nil", item)
	}
}

func TestBulkReadCache_Invalidate(t *testing.T) {
	c := NewBulkReadCache()
	primeCache(t, c, "/ip/dns/static", []MikrotikItem{{".id": "*1"}})
	c.Invalidate("/ip/dns/static")
	_, _, pathCached := c.Lookup("/ip/dns/static", "*1")
	if pathCached {
		t.Errorf("pathCached = true after Invalidate, want false")
	}
}

func TestBulkReadCache_DoBulkFetch_CoalescesConcurrentCallers(t *testing.T) {
	c := NewBulkReadCache()
	const workers = 50
	var fetchCount int64
	response := []MikrotikItem{{".id": "*1", "name": "only"}}

	fetchStarted := make(chan struct{})
	release := make(chan struct{})
	fetch := func() ([]MikrotikItem, error) {
		atomic.AddInt64(&fetchCount, 1)
		close(fetchStarted)
		<-release
		return response, nil
	}

	var wg sync.WaitGroup
	results := make([]MikrotikItem, workers)

	// Worker 0 starts the fetch and blocks inside it.
	wg.Add(1)
	go func() {
		defer wg.Done()
		items, err := c.DoBulkFetch("/ip/dns/static", fetch)
		if err != nil {
			t.Errorf("worker 0: %v", err)
			return
		}
		results[0] = items[0]
	}()

	<-fetchStarted

	wg.Add(workers - 1)
	for i := 1; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			items, err := c.DoBulkFetch("/ip/dns/static", fetch)
			if err != nil {
				t.Errorf("worker %d: %v", i, err)
				return
			}
			results[i] = items[0]
		}(i)
	}

	// Give the late workers time to enter DoBulkFetch before we release.
	// The singleflight group only coalesces callers that arrive while the
	// in-flight call is still running.
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()

	if got := atomic.LoadInt64(&fetchCount); got != 1 {
		t.Errorf("fetch invoked %d times, want 1 (singleflight must coalesce)", got)
	}
	for i, item := range results {
		if item["name"] != "only" {
			t.Errorf("worker %d got %v, want {name:only}", i, item)
		}
	}
	if _, found, _ := c.Lookup("/ip/dns/static", "*1"); !found {
		t.Errorf("cache was not populated after DoBulkFetch")
	}
}

// TestBulkReadCache_DoBulkFetch_InvalidateMidFetch_DiscardsStaleSnapshot
// regression-guards the cache/invalidation race: if Invalidate runs while a
// bulk fetch is in flight, the fetched (pre-invalidate) items must not be
// committed to the cache, otherwise concurrent writers would later read a
// snapshot that predates their own writes and see their new ids as missing.
func TestBulkReadCache_DoBulkFetch_InvalidateMidFetch_DiscardsStaleSnapshot(t *testing.T) {
	c := NewBulkReadCache()
	const path = "/ip/dns/static"

	// First fetch: starts, we Invalidate() before it returns, then it returns
	// the stale snapshot. populateIfFresh must reject it; DoBulkFetch must
	// observe errStaleBulkFetch internally and retry at the new generation.
	var attempts int64
	items, err := c.DoBulkFetch(path, func() ([]MikrotikItem, error) {
		n := atomic.AddInt64(&attempts, 1)
		if n == 1 {
			// Simulate a concurrent writer bumping the generation mid-fetch.
			c.Invalidate(path)
			return []MikrotikItem{{".id": "*stale"}}, nil
		}
		// The retry observes the post-invalidate generation and fetches fresh.
		return []MikrotikItem{{".id": "*fresh-1"}, {".id": "*fresh-2"}}, nil
	})
	if err != nil {
		t.Fatalf("DoBulkFetch err = %v", err)
	}
	if atomic.LoadInt64(&attempts) != 2 {
		t.Errorf("fetch attempts = %d, want 2 (first stale, retry fresh)", attempts)
	}
	if len(items) != 2 || items[0][".id"] != "*fresh-1" {
		t.Errorf("items = %v, want the fresh post-invalidate snapshot", items)
	}
	// Cache must hold only the fresh items; the stale *stale id must not leak.
	if _, found, _ := c.Lookup(path, "*stale"); found {
		t.Errorf("cache holds stale *stale — populate ran despite Invalidate")
	}
	if _, found, _ := c.Lookup(path, "*fresh-1"); !found {
		t.Errorf("cache missing *fresh-1 — retry did not populate")
	}
}

// TestBulkReadCache_DoBulkFetch_ConcurrentInvalidateDoesNotStrand
// exercises the worst case for singleflight coalescing: a bulk fetch is
// in-flight when a writer invalidates and then reads its own new id. The
// writer must not get the pre-write snapshot that the in-flight fetch would
// otherwise produce.
func TestBulkReadCache_DoBulkFetch_ConcurrentInvalidateDoesNotStrand(t *testing.T) {
	c := NewBulkReadCache()
	const path = "/ip/dns/static"

	release := make(chan struct{})
	firstStarted := make(chan struct{})
	var once sync.Once
	var attempts int64

	fetch := func() ([]MikrotikItem, error) {
		n := atomic.AddInt64(&attempts, 1)
		if n == 1 {
			once.Do(func() { close(firstStarted) })
			<-release // block first fetch until the writer arrives
			return []MikrotikItem{{".id": "*old"}}, nil
		}
		return []MikrotikItem{{".id": "*old"}, {".id": "*new"}}, nil
	}

	var readerItems []MikrotikItem
	var readerErr error
	readerDone := make(chan struct{})
	go func() {
		defer close(readerDone)
		readerItems, readerErr = c.DoBulkFetch(path, fetch)
	}()

	<-firstStarted
	// Writer invalidates (as a real CreateItem would after a successful write)
	// then issues its post-write read. This must not be served the pre-write
	// snapshot from the in-flight reader's fetch.
	c.Invalidate(path)
	var writerItems []MikrotikItem
	var writerErr error
	writerDone := make(chan struct{})
	go func() {
		defer close(writerDone)
		writerItems, writerErr = c.DoBulkFetch(path, fetch)
	}()

	// Give the writer time to enter DoBulkFetch and pick its generation before
	// we unblock the original (stale) fetch.
	time.Sleep(50 * time.Millisecond)
	close(release)

	<-readerDone
	<-writerDone

	if readerErr != nil {
		t.Fatalf("reader err = %v", readerErr)
	}
	if writerErr != nil {
		t.Fatalf("writer err = %v", writerErr)
	}
	// Both callers must ultimately see the post-invalidate snapshot that
	// contains *new. The writer would silently miss its own resource otherwise.
	for _, items := range [][]MikrotikItem{readerItems, writerItems} {
		var sawNew bool
		for _, it := range items {
			if it[".id"] == "*new" {
				sawNew = true
			}
		}
		if !sawNew {
			t.Errorf("got %v, want a snapshot containing *new (post-invalidate state)", items)
		}
	}
}

func TestBulkReadCache_DoBulkFetch_ErrorDoesNotPopulate(t *testing.T) {
	c := NewBulkReadCache()
	wantErr := fmt.Errorf("network boom")
	_, err := c.DoBulkFetch("/p", func() ([]MikrotikItem, error) { return nil, wantErr })
	if err != wantErr {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if _, _, pathCached := c.Lookup("/p", "*1"); pathCached {
		t.Errorf("cache was populated after failed fetch")
	}
}
