package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingBgpVpn = "routeros_routing_bgp_vpn.test"

func TestAccRoutingBgpVpnTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/vpn", "routeros_routing_bgp_vpn"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingBgpVpnConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBgpVpn),
							resource.TestCheckResourceAttr(testRoutingBgpVpn, "disabled", "false"),
							resource.TestCheckResourceAttr(testRoutingBgpVpn, "label_allocation_policy", "per-vrf"),
							resource.TestCheckResourceAttr(testRoutingBgpVpn, "name", "bgp-mpls-vpn-test"),
							resource.TestCheckResourceAttr(testRoutingBgpVpn, "route_distinguisher", "1.2.3.4:1"),
							resource.TestCheckResourceAttr(testRoutingBgpVpn, "vrf", "main")),
					},
				},
			})

		})
	}
}

func testAccRoutingBgpVpnConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bgp_vpn" "test" {
  disabled = false
  export {
    redistribute  = "connected"
    route_targets = ["1:1"]
  }
  import {
    route_targets = ["1:2"]
  }
  label_allocation_policy = "per-vrf"
  name                    = "bgp-mpls-vpn-test"
  route_distinguisher     = "1.2.3.4:1"
  vrf                     = "main"
}
`, providerConfig)
}
