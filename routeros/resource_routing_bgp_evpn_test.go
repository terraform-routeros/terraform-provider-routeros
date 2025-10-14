package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testBgpEvpnMinVersion = "7.20"
const testBgpEvpn = "routeros_routing_bgp_evpn.test"

func TestAccBgpEvpnTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testBgpEvpnMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testBgpEvpnMinVersion)
		return
	}

	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/evpn", "routeros_routing_bgp_evpn"),
				Steps: []resource.TestStep{
					{
						Config: testAccBgpEvpnConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testBgpEvpn),
							resource.TestCheckResourceAttr(testBgpEvpn, "name", "bgp-evpn-1"),
							resource.TestCheckResourceAttr(testBgpEvpn, "instance", "bgp-instance-1"),
							resource.TestCheckResourceAttr(testBgpEvpn, "vni", "1010"),
						),
					},
				},
			})

		})
	}
}

func testAccBgpEvpnConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_routing_bgp_instance" "test" {
  as   = "65000"
  name = "bgp-instance-1"
}

resource "routeros_routing_bgp_evpn" "test" {
  disabled = false
  export {
    route_targets = ["1010:1010"]
  }
  import {
    route_targets = ["1010:1010"]
  }
  instance = resource.routeros_routing_bgp_instance.test.name
  name     = "bgp-evpn-1"
  vni      = 1010
}
`, providerConfig)
}
