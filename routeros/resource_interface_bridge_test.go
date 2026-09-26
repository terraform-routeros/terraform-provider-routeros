package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeAddress = "routeros_interface_bridge.test_bridge"
const testInterfaceBridgeAddressWithSpace = "routeros_interface_bridge.test_bridge_w_space"

func TestAccInterfaceBridgeTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bridge", "routeros_interface_bridge"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddress, "name", "test_bridge"),
						),
					},
					{
						Config: testAccInterfaceBridgeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeAddressWithSpace),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddressWithSpace, "name", "Main bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceBridgeConfig() string {
	return providerConfig + `

resource "routeros_interface_bridge" "test_bridge" {
	name   = "test_bridge"
}

resource "routeros_interface_bridge" "test_bridge_w_space" {
	name              = "Main bridge"
	ageing_time       = "300s"
	ingress_filtering = true
	protocol_mode     = "rstp"
	priority          = "0x3000"
	igmp_snooping     = true
	vlan_filtering    = true
}

`
}

// RouterOS 7.22 adds these fields to bridge print replies (issue #965).
func TestInterfaceBridgeRouterOS722Read(t *testing.T) {
	originalVersion := RouterOSVersion
	RouterOSVersion = "7.22"
	t.Cleanup(func() { RouterOSVersion = originalVersion })
	res := ResourceInterfaceBridge()
	data := res.TestResourceData()
	item := MikrotikItem{
		".id": "*1", "name": "bridge1", "managed": "false",
		"mlag-peer-port": "none", "mlag-heartbeat": "5s",
		"mlag-priority": "128", "ra-guard": "false",
	}
	if diags := MikrotikResourceDataToTerraform(item, res.Schema, data); len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	for key, want := range map[string]interface{}{
		"managed": false, "mlag_peer_port": "none", "mlag_heartbeat": "5s",
		"mlag_priority": 128, "ra_guard": false,
	} {
		if got := data.Get(key); got != want {
			t.Errorf("%s = %#v, want %#v", key, got, want)
		}
	}
}

func TestInterfaceBridgeRouterOS722Write(t *testing.T) {
	originalVersion := RouterOSVersion
	RouterOSVersion = "7.22"
	t.Cleanup(func() { RouterOSVersion = originalVersion })
	res := ResourceInterfaceBridge()
	for _, config := range []map[string]interface{}{
		{"name": "bridge1"}, // Older RouterOS must not receive the new options unless configured.
		{"name": "bridge1", "mlag_peer_port": "ether2", "mlag_heartbeat": "1s", "mlag_priority": 0, "ra_guard": true},
		{"name": "bridge1", "mlag_peer_port": "none", "mlag_heartbeat": "none", "mlag_priority": 128, "ra_guard": false},
	} {
		data := testResourceDataWithConfig(t, res, config)
		if err := data.Set("managed", true); err != nil {
			t.Fatal(err)
		}
		item, _ := TerraformResourceDataToMikrotik(res.Schema, data)
		if _, ok := item["managed"]; ok {
			t.Error("read-only managed status must not be sent to RouterOS")
		}
		for _, key := range []string{"mlag_peer_port", "mlag_heartbeat", "mlag_priority", "ra_guard"} {
			value, configured := config[key]
			got, sent := item[SnakeToKebab(key)]
			if sent != configured {
				t.Errorf("%s: sent=%v, configured=%v", key, sent, configured)
				continue
			}
			if !configured {
				continue
			}
			want := fmt.Sprint(value)
			if b, ok := value.(bool); ok {
				want = "no"
				if b {
					want = "yes"
				}
			}
			if got != want {
				t.Errorf("%s = %q, want %q", key, got, want)
			}
		}
	}
}
