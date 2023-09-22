package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBondingAddress = "routeros_interface_bonding.test"

func TestAccInterfaceBondingTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bonding", "routeros_interface_bonding"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBondingConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBondingAddress),
							resource.TestCheckResourceAttr(testInterfaceBondingAddress, "name", "bonding-test"),
						),
					},
				},
			})
		})
	}
}

func testAccInterfaceBondingConfig() string {
	return providerConfig + `

resource "routeros_interface_bonding" "test" {
  name   = "bonding-test"
  slaves =  ["ether3", "ether4"]
}
`
}
