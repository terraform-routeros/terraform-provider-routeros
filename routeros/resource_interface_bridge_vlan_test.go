package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceBridgeVlanAddress = "routeros_bridge_vlan.test_vlan"

func TestAccInterfaceBridgeVlanTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/vlan", "routeros_bridge_vlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceBridgeVlanExists(testInterfaceBridgeVlanAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeVlanAddress, "bridge", "bridge"),
						),
					},
				},
			})

		})
	}

}

func testAccCheckInterfaceBridgeVlanExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccInterfaceBridgeVlanConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_bridge_vlan" "test_vlan" {
	bridge   = "bridge"
	untagged = ["ether1"]
	tagged	 = ["bridge"]
	vlan_ids = 200
  }

`
}
