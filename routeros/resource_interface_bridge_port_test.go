package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceBridgePortAddress = "routeros_interface_bridge_port.test_port"

func TestAccInterfaceBridgePortTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bridge/port", "routeros_interface_bridge_port"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgePortConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceBridgePortExists(testInterfaceBridgePortAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgePortAddress, "bridge", "bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceBridgePortExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceBridgePortConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_bridge_port" "test_port" {
	bridge    = "bridge"
	interface = "ether1"
	pvid 	  = 200
	disabled  = true
  }

`
}
