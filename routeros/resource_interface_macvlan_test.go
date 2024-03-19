package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceMacVlanAddress = "routeros_interface_macvlan.test"

func TestAccInterfaceMacVlanTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/macvlan", "routeros_interface_macvlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceMacVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMacVlanAddress),
							resource.TestCheckResourceAttr(testInterfaceMacVlanAddress, "name", "macvlan1"),
						),
					},
				},
			})
		})
	}
}

func testAccInterfaceMacVlanConfig() string {
	return providerConfig + `

resource "routeros_interface_macvlan" "test" {
  interface    = "ether1" 
  name         = "macvlan1"
  disabled     = false
}
`
}
