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

func TestBulkReadCache_PopulateAndLookup_Hit(t *testing.T) {
	c := NewBulkReadCache()
	c.populate("/ip/dns/static", []MikrotikItem{
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
	c.populate("/ip/dns/static", []MikrotikItem{{".id": "*1"}})
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
	c.populate("/ip/dns/static", []MikrotikItem{{".id": "*1"}})
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
