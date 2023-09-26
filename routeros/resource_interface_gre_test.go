package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testGreAddress = "routeros_interface_gre.gre900"
const testGreName = "GRE_900_TEST"

func TestAccInterfaceGreTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/gre", "routeros_interface_gre"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceGreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testGreAddress),
							resource.TestCheckResourceAttr(testGreAddress, "name", testGreName),
						),
					},
				},
			})

		})
	}

}

func testAccInterfaceGreConfig() string {
	return providerConfig + `

resource "routeros_interface_gre" "gre900" {
	name      = "GRE_900_TEST"
	remote_address = "127.0.0.1"
	disabled  = true
}
`
}
