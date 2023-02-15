package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
							testAccCheckInterfaceGreExists(testGreAddress),
							resource.TestCheckResourceAttr(testGreAddress, "name", testGreName),
						),
					},
				},
			})

		})
	}

}

func testAccCheckInterfaceGreExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceGreConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_gre" "gre900" {
	name      = "GRE_900_TEST"
	remote_address = "127.0.0.1"
	disabled  = true
}
`
}
