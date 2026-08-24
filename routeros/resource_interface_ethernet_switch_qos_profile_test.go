package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Hardware validation (manual, real switches):
//   - Marvell-98DX8216 / CRS317-1G-16S+ / RouterOS 7.20.8: full create -> update -> delete,
//     each step verified via the REST API; switch config restored to its pre-test baseline.
//   - Marvell-98DX226S / CRS310-8G+2S+ / RouterOS 7.16.2: create + delete of an inert profile.
//   - Marvell-98DX8208 / CRS309-1G-8S+ / RouterOS 7.18.2: import + plan round-trip (no diff).
// The test below encodes the same create/update/delete flow. It targets DSCP 63, which is
// unused in a default DSCP layout, so running it against a live switch does not affect traffic.

const testInterfaceEthernetSwitchQosProfile = "routeros_interface_ethernet_switch_qos_profile.test"

func TestAccInterfaceEthernetSwitchQosProfileTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ethernet/switch/qos/profile", "routeros_interface_ethernet_switch_qos_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceEthernetSwitchQosProfileConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceEthernetSwitchQosProfile),
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosProfile, "name", "tf-test-inert"),
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosProfile, "dscp", "63"),
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosProfile, "traffic_class", "0"),
						),
					},
					{
						Config: testAccInterfaceEthernetSwitchQosProfileConfig(2),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosProfile, "traffic_class", "2"),
						),
					},
				},
			})
		})
	}
}

func testAccInterfaceEthernetSwitchQosProfileConfig(trafficClass int) string {
	return fmt.Sprintf(`%v

resource "routeros_interface_ethernet_switch_qos_profile" "test" {
  name          = "tf-test-inert"
  dscp          = 63
  traffic_class = %v
}
`, providerConfig, trafficClass)
}
