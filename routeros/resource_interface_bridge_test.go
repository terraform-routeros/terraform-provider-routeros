package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckInterfaceBridgeExists(testInterfaceBridgeAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddress, "name", "test_bridge"),
						),
					},
					{
						Config: testAccInterfaceBridgeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceBridgeExists(testInterfaceBridgeAddressWithSpace),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddressWithSpace, "name", "Main bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceBridgeExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceBridgeConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_bridge" "test_bridge" {
	name   = "test_bridge"
}

resource "routeros_interface_bridge" "test_bridge_w_space" {
	name   = "Main bridge"
	ingress_filtering = true
	protocol_mode = "rstp"
	priority = "0x3000"
	igmp_snooping = true
	vlan_filtering = true
}

`
}
