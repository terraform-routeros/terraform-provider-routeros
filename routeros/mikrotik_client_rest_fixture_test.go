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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func TestRestClientVlanRequestSendsExactL3HWValues(t *testing.T) {
	var requestMethods []string
	var requestBodies []map[string]string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody map[string]string
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Fatal(err)
		}
		requestMethods = append(requestMethods, r.Method)
		requestBodies = append(requestBodies, requestBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ret":".id=*1"}`)
	}))
	defer server.Close()

	client := &RestClient{
		ctx:       context.Background(),
		HostURL:   server.URL,
		Transport: TransportREST,
		Client:    server.Client(),
	}

	_, err := CreateItem(
		context.Background(),
		MikrotikItem{
			"name":             "VLAN900",
			"interface":        "bridge",
			"vlan-id":          "900",
			"l3-hw-offloading": "yes",
		},
		"/interface/vlan",
		client,
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = UpdateItem(
		&ItemId{Id, "*1"},
		"/interface/vlan",
		MikrotikItem{
			"name":             "VLAN900",
			"interface":        "bridge",
			"vlan-id":          "900",
			"l3-hw-offloading": "no",
		},
		client,
	)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := requestMethods, []string{http.MethodPut, http.MethodPatch}; !reflect.DeepEqual(got, want) {
		t.Fatalf("request methods = %#v, want %#v", got, want)
	}
	for i, expected := range []string{"yes", "no"} {
		if got := requestBodies[i]["l3-hw-offloading"]; got != expected {
			t.Fatalf("request %d l3-hw-offloading = %q, want %q", i, got, expected)
		}
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

func TestIPAddressUpdatePreservesStateAfterRejectedPatch(t *testing.T) {
	RouterOSVersion = "7.24.5"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			_, _ = io.WriteString(w, `[{"ret":".id=*1","address":"192.0.2.1/32","interface":"vlan10","vrf":"main"}]`)
			return
		}
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
	resourceSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/address"),
		MetaId:           PropId(Id),
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"interface": {
			Type:     schema.TypeString,
			Optional: true,
		},
		KeyVrf: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	resourceData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
		"address":   "192.0.2.1/32",
		"interface": "vlan10",
		"vrf":       "main",
	})
	resourceData.SetId("*1")

	diags := UpdateIPAddress(resourceSchema)(context.Background(), resourceData, client)
	if !diags.HasError() {
		t.Fatal("expected rejected PATCH diagnostic")
	}
	if got := resourceData.Get("vrf"); got != "main" {
		t.Fatalf("vrf state after rejected PATCH = %q, want %q", got, "main")
	}
}
