package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceBridgeAddress = "routeros_bridge.test_bridge"

func TestAccInterfaceBridgeTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge", "routeros_bridge"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceBridgeExists(testInterfaceBridgeAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddress, "name", "test_bridge"),
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

resource "routeros_bridge" "test_bridge" {
	name   = "test_bridge"
  }

`
}
