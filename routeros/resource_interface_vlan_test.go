package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testVlanAddress = "routeros_vlan.vlan900"
const testVlanName = "VLAN_900_TEST"

func TestAccInterfaceVlanTest(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/interface/vlan", "routeros_vlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceVlanExists(testVlanAddress),
							resource.TestCheckResourceAttr(testVlanAddress, "name", testVlanName),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceVlanExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceVlanConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_vlan" "vlan900" {
	name      = "VLAN_900_TEST"
	vlan_id   = 900
	disabled  = true
	interface = "bridge"
}
`
}
