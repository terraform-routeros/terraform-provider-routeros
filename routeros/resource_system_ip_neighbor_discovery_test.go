package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpNeighborDiscoverySettings = "routeros_ip_neighbor_discovery_settings.test"

func TestAccIpNeighborDiscoverySettingsTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccIpNeighborDiscoverySettingsConfig("static", "1", "rx-only", `[]`),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpNeighborDiscoverySettings),
							resource.TestCheckResourceAttr(testIpNeighborDiscoverySettings, "discover_interface_list", "static"),
						),
					},
					{
						Config: testAccIpNeighborDiscoverySettingsConfig("LAN", "disabled", "tx-and-rx", `["cdp", "lldp", "mndp"]`),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testIpNeighborDiscoverySettings, "discover_interface_list", "LAN"),
						),
					},
				},
			})
		})
	}
}

func testAccIpNeighborDiscoverySettingsConfig(iflist, lldp, mode, proto string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_neighbor_discovery_settings" "test" {
	discover_interface_list  = "%v"
	lldp_med_net_policy_vlan = "%v"
	mode                     = "%v"
	protocol                 =  %v
}
`, providerConfig, iflist, lldp, mode, proto)
}
