package routeros

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUpdateDhcpServerRejectsRESTResponseThatOmitsConfiguredFields(t *testing.T) {
	RouterOSVersion = "7.24.5"
	var patchBody map[string]string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPatch {
			t.Fatalf("unexpected method %s", r.Method)
		}
		if err := json.NewDecoder(r.Body).Decode(&patchBody); err != nil {
			t.Fatal(err)
		}
		_, _ = io.WriteString(w, `{"ret":".id=*1","name":"dhcp1"}`)
	}))
	defer server.Close()

	client := &RestClient{
		ctx:       context.Background(),
		HostURL:   server.URL,
		Transport: TransportREST,
		Client:    server.Client(),
	}
	resourceSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server"),
		MetaId:           PropId(Id),
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		dhcpServerAuthoritative: {
			Type:     schema.TypeString,
			Optional: true,
		},
		dhcpServerBootpLease: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	resourceData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
		"name":                  "dhcp1",
		dhcpServerAuthoritative: "yes",
		dhcpServerBootpLease:    "forever",
	})
	res, err := UpdateItem(&ItemId{Id, "*1"}, "/ip/dhcp-server", MikrotikItem{
		"authoritative":    "yes",
		"bootp-lease-time": "forever",
	}, client)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateDhcpServerRESTFields(resourceData, res, client); err == nil {
		t.Fatal("expected an error when RouterOS omits configured DHCP fields")
	} else if !strings.Contains(err.Error(), "authoritative") {
		t.Fatalf("error = %q, want omitted field name", err)
	}
	if got := patchBody[SnakeToKebab(dhcpServerAuthoritative)]; got != "yes" {
		t.Fatalf("authoritative PATCH value = %q, want %q", got, "yes")
	}
	if got := patchBody[SnakeToKebab(dhcpServerBootpLease)]; got != "forever" {
		t.Fatalf("bootp lease PATCH value = %q, want %q", got, "forever")
	}
}
