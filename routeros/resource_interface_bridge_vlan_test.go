package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeVlanAddress = "routeros_interface_bridge_vlan.test_vlan"

func TestAccInterfaceBridgeVlanTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bridge/vlan", "routeros_interface_bridge_vlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeVlanAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeVlanAddress, "bridge", "bridge"),
						),
					},
				},
			})

		})
	}

}

func testAccInterfaceBridgeVlanConfig() string {
	return providerConfig + `

resource "routeros_interface_bridge_vlan" "test_vlan" {
	bridge   = "bridge"
	untagged = ["ether1"]
	tagged	 = ["bridge"]
	vlan_ids = [200]
  }

`
}
