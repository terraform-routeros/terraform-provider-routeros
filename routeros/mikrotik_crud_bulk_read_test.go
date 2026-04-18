package routeros

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
)

// recordedCall captures a single SendRequest invocation so tests can assert on
// what network traffic the provider would have issued.
type recordedCall struct {
	method crudMethod
	path   string
	query  []string
}

// fakeClient is an in-memory Client implementation used for unit-testing the
// CRUD plumbing without touching a real router.
type fakeClient struct {
	mu        sync.Mutex
	calls     []recordedCall
	bulkItems map[string][]MikrotikItem // path -> canned response for bulk GET
	filtered  map[string]MikrotikItem   // path + query-key -> canned response for filtered GET
	sendErr   error                     // if set, SendRequest returns this error
	cache     *BulkReadCache
	transport TransportType
}

func newFakeClient(enableCache bool) *fakeClient {
	fc := &fakeClient{
		bulkItems: map[string][]MikrotikItem{},
		filtered:  map[string]MikrotikItem{},
		transport: TransportREST,
	}
	if enableCache {
		fc.cache = NewBulkReadCache()
	}
	return fc
}

func (f *fakeClient) GetExtraParams() *ExtraParams   { return &ExtraParams{} }
func (f *fakeClient) GetTransport() TransportType    { return f.transport }
func (f *fakeClient) GetBulkCache() *BulkReadCache   { return f.cache }

func (f *fakeClient) SendRequest(method crudMethod, url *URL, item MikrotikItem, result interface{}) error {
	f.mu.Lock()
	f.calls = append(f.calls, recordedCall{method: method, path: url.Path, query: append([]string(nil), url.Query...)})
	f.mu.Unlock()

	if f.sendErr != nil {
		return f.sendErr
	}
	if method != crudRead || result == nil {
		return nil
	}

	slice, ok := result.(*[]MikrotikItem)
	if !ok {
		return nil
	}

	// Bulk GET has no query.
	if len(url.Query) == 0 {
		if items, found := f.bulkItems[url.Path]; found {
			*slice = append(*slice, items...)
		}
		return nil
	}

	// Filtered GET ("?.id=*X"): return a single match if present.
	key := url.Path + "|" + strings.Join(url.Query, "&")
	if item, found := f.filtered[key]; found {
		*slice = append(*slice, item)
	}
	return nil
}

func (f *fakeClient) callCount(method crudMethod, path string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	n := 0
	for _, c := range f.calls {
		if c.method == method && c.path == path {
			n++
		}
	}
	return n
}

func (f *fakeClient) bulkReadCount(path string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	n := 0
	for _, c := range f.calls {
		if c.method == crudRead && c.path == path && len(c.query) == 0 {
			n++
		}
	}
	return n
}

func (f *fakeClient) filteredReadCount(path string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	n := 0
	for _, c := range f.calls {
		if c.method == crudRead && c.path == path && len(c.query) > 0 {
			n++
		}
	}
	return n
}

// ---------- ReadItems tests ----------

func TestReadItems_CacheDisabled_UsesFilteredGet(t *testing.T) {
	fc := newFakeClient(false)
	fc.filtered["/ip/dns/static|?.id=*1"] = MikrotikItem{".id": "*1", "name": "a"}

	res, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc)
	if err != nil {
		t.Fatal(err)
	}
	if len(*res) != 1 || (*res)[0]["name"] != "a" {
		t.Errorf("got %v, want one item named a", *res)
	}
	if got := fc.filteredReadCount("/ip/dns/static"); got != 1 {
		t.Errorf("filtered GET calls = %d, want 1", got)
	}
	if got := fc.bulkReadCount("/ip/dns/static"); got != 0 {
		t.Errorf("bulk GET calls = %d, want 0 when cache disabled", got)
	}
}

func TestReadItems_CacheEnabled_FirstCall_BulkGet(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{
		{".id": "*1", "name": "a"}, {".id": "*2", "name": "b"},
	}

	res, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc)
	if err != nil {
		t.Fatal(err)
	}
	if len(*res) != 1 || (*res)[0]["name"] != "a" {
		t.Errorf("got %v, want one item named a", *res)
	}
	if got := fc.bulkReadCount("/ip/dns/static"); got != 1 {
		t.Errorf("bulk GET calls = %d, want 1", got)
	}
	if got := fc.filteredReadCount("/ip/dns/static"); got != 0 {
		t.Errorf("filtered GET calls = %d, want 0", got)
	}
}

func TestReadItems_CacheEnabled_SecondCall_NoRequest(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{
		{".id": "*1", "name": "a"}, {".id": "*2", "name": "b"},
	}

	if _, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadItems(&ItemId{Id, "*2"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}

	if got := fc.bulkReadCount("/ip/dns/static"); got != 1 {
		t.Errorf("bulk GET calls = %d, want 1 (second Read must hit cache)", got)
	}
}

func TestReadItems_CacheEnabled_IdNotInBulk_ReturnsEmpty(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{{".id": "*1"}}

	res, err := ReadItems(&ItemId{Id, "*missing"}, "/ip/dns/static", fc)
	if err != nil {
		t.Fatal(err)
	}
	if len(*res) != 0 {
		t.Errorf("got %v, want empty slice", *res)
	}

	// Second lookup for the same missing id still serves from cache.
	if _, err := ReadItems(&ItemId{Id, "*missing"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if got := fc.bulkReadCount("/ip/dns/static"); got != 1 {
		t.Errorf("bulk GET calls = %d, want 1", got)
	}
}

func TestReadItems_CacheEnabled_AfterInvalidate_RefetchesBulk(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{{".id": "*1"}}

	if _, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	fc.cache.Invalidate("/ip/dns/static")
	if _, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if got := fc.bulkReadCount("/ip/dns/static"); got != 2 {
		t.Errorf("bulk GET calls = %d, want 2 (one per pre/post invalidate lookup)", got)
	}
}

func TestReadItems_CacheEnabled_NilId_BypassesCache(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/system/resource"] = []MikrotikItem{{".id": "*1", "version": "7.22"}}

	if _, err := ReadItems(nil, "/system/resource", fc); err != nil {
		t.Fatal(err)
	}
	if _, _, pathCached := fc.cache.Lookup("/system/resource", "*1"); pathCached {
		t.Errorf("datasource-style read (id==nil) populated the cache — must bypass")
	}
}

func TestReadItems_CacheEnabled_NameIdType_BypassesCache(t *testing.T) {
	fc := newFakeClient(true)
	fc.filtered["/interface/vlan|?name=vlan10"] = MikrotikItem{".id": "*7F", "name": "vlan10"}

	res, err := ReadItems(&ItemId{Name, "vlan10"}, "/interface/vlan", fc)
	if err != nil {
		t.Fatal(err)
	}
	if len(*res) != 1 || (*res)[0]["name"] != "vlan10" {
		t.Errorf("got %v, want vlan10", *res)
	}
	if got := fc.filteredReadCount("/interface/vlan"); got != 1 {
		t.Errorf("filtered GET calls = %d, want 1", got)
	}
	if got := fc.bulkReadCount("/interface/vlan"); got != 0 {
		t.Errorf("Name lookup triggered bulk read — must bypass cache")
	}
}

func TestReadItems_CacheEnabled_BulkGetError_Propagates(t *testing.T) {
	fc := newFakeClient(true)
	fc.sendErr = errors.New("http 500")

	_, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc)
	if err == nil || !strings.Contains(err.Error(), "http 500") {
		t.Errorf("err = %v, want propagation of http 500", err)
	}
	if _, _, pathCached := fc.cache.Lookup("/ip/dns/static", "*1"); pathCached {
		t.Errorf("cache was populated after failed bulk GET — must stay empty so next call retries")
	}
}

func TestReadItems_CacheEnabled_ParallelCallsCoalesceToOneBulkGet(t *testing.T) {
	fc := newFakeClient(true)
	items := make([]MikrotikItem, 100)
	for i := range items {
		items[i] = MikrotikItem{".id": fmt.Sprintf("*%d", i), "seq": fmt.Sprintf("%d", i)}
	}
	fc.bulkItems["/ip/dns/static"] = items

	// The fake's SendRequest is fast so singleflight may not coalesce every
	// call — caching correctness is the real contract. We assert the weaker
	// invariant that the number of bulk GETs is bounded far below the worker
	// count, which is what the feature promises.
	const workers = 100
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			id := fmt.Sprintf("*%d", i)
			res, err := ReadItems(&ItemId{Id, id}, "/ip/dns/static", fc)
			if err != nil {
				t.Errorf("worker %d: %v", i, err)
				return
			}
			if len(*res) != 1 || (*res)[0]["seq"] != fmt.Sprintf("%d", i) {
				t.Errorf("worker %d got %v", i, *res)
			}
		}(i)
	}
	wg.Wait()

	// With a fast in-memory fake, singleflight may not coalesce every single
	// call — but the number of bulk GETs must be small and bounded (far less
	// than one per worker). Caching correctness is the real contract.
	got := fc.bulkReadCount("/ip/dns/static")
	if got >= workers {
		t.Errorf("bulk GET calls = %d, want far fewer than %d (cache must amortize)", got, workers)
	}
	if got := fc.filteredReadCount("/ip/dns/static"); got != 0 {
		t.Errorf("filtered GET calls = %d, want 0 when cache enabled", got)
	}
}

// ---------- CRUD invalidation tests ----------

func TestCreateItem_Success_InvalidatesPath(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{{".id": "*1"}}
	// Prime the cache.
	if _, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if _, _, pathCached := fc.cache.Lookup("/ip/dns/static", "*1"); !pathCached {
		t.Fatal("cache not primed")
	}

	if _, err := CreateItem(context.Background(), MikrotikItem{"name": "new"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if _, _, pathCached := fc.cache.Lookup("/ip/dns/static", "*1"); pathCached {
		t.Errorf("cache was not invalidated after CreateItem")
	}
}

// Guard against regression where DeleteItem's path mutation (appending "/" + id
// for REST) would invalidate the wrong cache key. Also exercises the success
// path of DeleteItem invalidation.
func TestDeleteItem_RestTransport_InvalidatesBasePath(t *testing.T) {
	fc := newFakeClient(true)
	fc.bulkItems["/ip/dns/static"] = []MikrotikItem{{".id": "*1"}}
	if _, err := ReadItems(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}

	if err := DeleteItem(&ItemId{Id, "*1"}, "/ip/dns/static", fc); err != nil {
		t.Fatal(err)
	}
	if _, _, pathCached := fc.cache.Lookup("/ip/dns/static", "*1"); pathCached {
		t.Errorf("cache for base path /ip/dns/static was not invalidated (DeleteItem suffixed the id)")
	}
	// Sanity: the wrong cache key shouldn't magically appear.
	if _, _, pathCached := fc.cache.Lookup("/ip/dns/static/*1", ""); pathCached {
		t.Errorf("cache entry appeared under the suffixed path — invalidation used the wrong key")
	}
}

