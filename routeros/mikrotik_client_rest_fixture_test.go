package routeros

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestRestClientUpdateItemSendsExactPatchBody(t *testing.T) {
	var requestMethod, requestPath string
	var requestBody map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestMethod = r.Method
		requestPath = r.URL.Path
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ret":".id=*1","comment":"updated"}`)
	}))
	defer server.Close()

	client := &RestClient{
		ctx:       context.Background(),
		HostURL:   server.URL,
		Transport: TransportREST,
		Client:    server.Client(),
	}

	_, err := UpdateItem(
		&ItemId{Id, "*1"},
		"/ip/address",
		MikrotikItem{"address": "192.0.2.1/32", "interface": "vlan10", "comment": "updated"},
		client,
	)
	if err != nil {
		t.Fatal(err)
	}

	if requestMethod != http.MethodPatch {
		t.Fatalf("method = %q, want %q", requestMethod, http.MethodPatch)
	}
	if requestPath != "/rest/ip/address/*1" {
		t.Fatalf("path = %q, want %q", requestPath, "/rest/ip/address/*1")
	}
	expected := map[string]string{
		"address":   "192.0.2.1/32",
		"interface": "vlan10",
		"comment":   "updated",
	}
	if !reflect.DeepEqual(requestBody, expected) {
		t.Fatalf("PATCH body = %#v, want %#v", requestBody, expected)
	}
}

func TestRestClientUpdateItemReturnsFourHundredError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, `{"error":400,"message":"Bad Request","detail":"unknown parameter vrf"}`)
	}))
	defer server.Close()

	client := &RestClient{
		ctx:       context.Background(),
		HostURL:   server.URL,
		Transport: TransportREST,
		Client:    server.Client(),
	}

	_, err := UpdateItem(
		&ItemId{Id, "*1"},
		"/ip/address",
		MikrotikItem{"address": "192.0.2.1/32", "interface": "vlan10"},
		client,
	)
	if err == nil {
		t.Fatal("expected PATCH failure for HTTP 400 response")
	}
	if got := err.Error(); got == "" || !strings.Contains(got, "unknown parameter vrf") {
		t.Fatalf("error = %q, want RouterOS error detail", got)
	}
}
