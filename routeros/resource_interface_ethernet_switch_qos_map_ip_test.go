package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Hardware validation (manual, real switches):
//   - Marvell-98DX8216 / CRS317-1G-16S+ / RouterOS 7.20.8: full create -> delete, verified via
//     the REST API; switch config restored to its pre-test baseline.
//   - Marvell-98DX226S / CRS310-8G+2S+ / RouterOS 7.16.2: create + delete of an inert map entry.
//   - Marvell-98DX8208 / CRS309-1G-8S+ / RouterOS 7.18.2: import + plan round-trip (no diff).
// The test creates a profile and a map entry for DSCP 63 (unused in a default DSCP layout), so
// running it against a live switch does not affect traffic.

const testInterfaceEthernetSwitchQosMapIp = "routeros_interface_ethernet_switch_qos_map_ip.test"

func TestAccInterfaceEthernetSwitchQosMapIpTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ethernet/switch/qos/map/ip", "routeros_interface_ethernet_switch_qos_map_ip"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceEthernetSwitchQosMapIpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceEthernetSwitchQosMapIp),
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosMapIp, "dscp", "63"),
							resource.TestCheckResourceAttr(testInterfaceEthernetSwitchQosMapIp, "profile", "tf-test-inert"),
						),
					},
				},
			})
		})
	}
}

func testAccInterfaceEthernetSwitchQosMapIpConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_ethernet_switch_qos_profile" "test" {
  name          = "tf-test-inert"
  dscp          = 63
  traffic_class = 0
}

resource "routeros_interface_ethernet_switch_qos_map_ip" "test" {
  dscp    = 63
  profile = routeros_interface_ethernet_switch_qos_profile.test.name
}
`, providerConfig)
}
