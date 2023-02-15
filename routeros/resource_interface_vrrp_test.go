package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceVrrpAddress = "routeros_interface_vrrp.test_vrrp_interface"

func TestAccInterfaceVrrpTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vrrp", "routeros_interface_vrrp"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVrrpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceVrrpExists(testInterfaceVrrpAddress),
							resource.TestCheckResourceAttr(testInterfaceVrrpAddress, "interface", "ether1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceVrrpExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceVrrpConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_vrrp" "test_vrrp_interface" {
	name   		= "test_vrrp_interface"
	interface = "ether1"
  }

`
}
