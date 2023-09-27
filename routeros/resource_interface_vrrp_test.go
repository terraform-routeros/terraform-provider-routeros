package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testInterfaceVrrpAddress),
							resource.TestCheckResourceAttr(testInterfaceVrrpAddress, "interface", "ether1"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVrrpConfig() string {
	return providerConfig + `
resource "routeros_interface_vrrp" "test_vrrp_interface" {
	name   		= "test_vrrp_interface"
	interface = "ether1"
  }

`
}
