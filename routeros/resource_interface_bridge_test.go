package routeros

import (
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
