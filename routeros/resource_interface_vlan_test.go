package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testVlanAddress = "routeros_interface_vlan.vlan900"
const testVlanName = "VLAN_900_TEST"

func TestAccInterfaceVlanTest(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vlan", "routeros_interface_vlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testVlanAddress),
							resource.TestCheckResourceAttr(testVlanAddress, "name", testVlanName),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVlanConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_vlan" "vlan900" {
	name      = "VLAN_900_TEST"
	vlan_id   = 900
	disabled  = true
	interface = "bridge"
}
`
}
