package routeros

import (
	"maps"
	"reflect"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/gocty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Supply RawConfig as well as state: the serializer distinguishes an omitted
// option from an explicitly configured zero or false value.
func testResourceDataWithConfig(t *testing.T, res *schema.Resource, config map[string]interface{}) *schema.ResourceData {
	t.Helper()
	raw := map[string]cty.Value{}
	for key, typ := range res.CoreConfigSchema().ImpliedType().AttributeTypes() {
		raw[key] = cty.NullVal(typ)
		if value, ok := config[key]; ok {
			var err error
			raw[key], err = gocty.ToCtyValue(value, typ)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	data := res.Data(&terraform.InstanceState{RawConfig: cty.ObjectVal(raw)})
	for key, value := range config {
		if err := data.Set(key, value); err != nil {
			t.Fatal(err)
		}
	}
	return data
}

func TestVersionedFieldConversions(t *testing.T) {
	originalVersion := RouterOSVersion
	t.Cleanup(func() { RouterOSVersion = originalVersion })
	tests := []struct {
		name, version string
		resource      func() *schema.Resource
		config        map[string]interface{}
		input, want   MikrotikItem
	}{
		{"legacy bridge", "7.22", ResourceInterfaceBridge, map[string]interface{}{"add_dhcp_option82": true}, MikrotikItem{"add-dhcp-option82": "yes"}, MikrotikItem{"add-dhcp-option82": "yes"}},
		{"upgraded bridge", "7.23", ResourceInterfaceBridge, map[string]interface{}{"add_dhcp_option82": true}, MikrotikItem{"add-dhcp-option82": "yes"}, MikrotikItem{"dhcp-agent-circuit-id": "$(INTERFACE):$(VID)", "dhcp-agent-remote-id": "$(BRIDGEMAC)"}},
		{"disable option82", "7.24.4", ResourceInterfaceBridge, map[string]interface{}{"add_dhcp_option82": false}, MikrotikItem{"add-dhcp-option82": "no", "dhcp-agent-circuit-id": "old", "dhcp-agent-remote-id": "old"}, MikrotikItem{"!dhcp-agent-circuit-id": "", "!dhcp-agent-remote-id": ""}},
		{"unconfigured option82", "7.24.4", ResourceInterfaceBridge, map[string]interface{}{}, MikrotikItem{"add-dhcp-option82": "no", "dhcp-agent-circuit-id": "custom"}, MikrotikItem{"dhcp-agent-circuit-id": "custom"}},
		{"legacy BGP", "7.21", ResourceRoutingBgpConnection, map[string]interface{}{"add_path_out": "all"}, MikrotikItem{"add-path-out": "all"}, MikrotikItem{"add-path-out": "all"}},
		{"upgraded BGP", "7.22", ResourceRoutingBgpConnection, map[string]interface{}{"add_path_out": "all"}, MikrotikItem{"add-path-out": "all"}, MikrotikItem{"output.add-path": "ip,ipv6"}},
		{"disable BGP addpath", "7.24.4", ResourceRoutingBgpTemplate, map[string]interface{}{"add_path_out": "none"}, MikrotikItem{"add-path-out": "none"}, MikrotikItem{"output.add-path": ""}},
		{"native BGP addpath", "7.24.4", ResourceRoutingBgpTemplate, map[string]interface{}{}, MikrotikItem{"add-path-out": "none", "output.add-path": "ipv6"}, MikrotikItem{"output.add-path": "ipv6"}},
		{"legacy port priority", "7.23", ResourceInterfaceBridgePort, map[string]interface{}{"priority": "80"}, MikrotikItem{"priority": "80"}, MikrotikItem{"priority": "80"}},
		{"hex port priority", "7.24.4", ResourceInterfaceBridgePort, map[string]interface{}{"priority": "80"}, MikrotikItem{"priority": "80"}, MikrotikItem{"priority": "0x50"}},
		{"hex port priority passthrough", "7.24.4", ResourceInterfaceBridgePort, map[string]interface{}{"priority": "0x80"}, MikrotikItem{"priority": "0x80"}, MikrotikItem{"priority": "0x80"}},
		{"legacy bsd-syslog", "7.23", ResourceSystemLoggingAction, map[string]interface{}{"bsd_syslog": true}, MikrotikItem{"bsd-syslog": "yes"}, MikrotikItem{"bsd-syslog": "yes"}},
		{"removed bsd-syslog", "7.24.4", ResourceSystemLoggingAction, map[string]interface{}{"bsd_syslog": true}, MikrotikItem{"bsd-syslog": "yes"}, MikrotikItem{}},
		{"deprecated cluster-id", "7.24.4", ResourceRoutingBgpConnection, map[string]interface{}{"cluster_id": "0.0.0.0", "router_id": "0.0.0.1"}, MikrotikItem{"cluster-id": "0.0.0.0", "router-id": "0.0.0.1", "input.ignore-as-path-len": "true"}, MikrotikItem{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RouterOSVersion = tt.version
			res := tt.resource()
			data := testResourceDataWithConfig(t, res, tt.config)
			got := maps.Clone(tt.input)
			versionedWriteFields(got, res.Schema[MetaResourcePath].Default.(string), data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %#v, want %#v", got, tt.want)
			}
			// A read after an upgrade must keep explicitly configured legacy values.
			if tt.name == "upgraded bridge" || tt.name == "upgraded BGP" {
				if diags := MikrotikResourceDataToTerraform(tt.want, res.Schema, data); len(diags) != 0 {
					t.Fatal(diags)
				}
				for key, want := range tt.config {
					if got := data.Get(key); got != want {
						t.Errorf("%s = %v, want %v", key, got, want)
					}
				}
			}
		})
	}
}

func TestRouterOS724RenamedFields(t *testing.T) {
	originalVersion := RouterOSVersion
	t.Cleanup(func() { RouterOSVersion = originalVersion })
	for _, tt := range []struct {
		res      *schema.Resource
		old, new string
	}{
		{ResourceIpService(), "address", "available-from"},
		{ResourceInterfaceDot1xServer(), "reauth_timeout", "reauth-period"},
		{DatasourceSystemResource(), "factory_software", "minimum-version"},
		{DatasourceSystemRouterboard(), "factory_firmware", "minimum-firmware"},
	} {
		t.Run(tt.old, func(t *testing.T) {
			RouterOSVersion = "7.24.4"
			data := tt.res.TestResourceData()
			if diags := MikrotikResourceDataToTerraform(MikrotikItem{tt.new: "test"}, tt.res.Schema, data); len(diags) != 0 {
				t.Fatal(diags)
			}
			if got := data.Get(tt.old); got != "test" {
				t.Fatalf("got %v", got)
			}
			RouterOSVersion = "7.23.7"
			path := tt.res.Schema[MetaResourcePath].Default.(string)
			if _, ok := driftAttributeSlice.GetDriftMap(RouterOSVersion, path, false)[tt.old]; ok {
				t.Fatal("rename applied before 7.24")
			}
		})
	}
}

func TestProviderSchemaCompatibility(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
