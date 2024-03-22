package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceMacVlanAddress = "routeros_interface_macvlan.test"

// resource introduced on 7.12.1 https://forum.mikrotik.com/viewtopic.php?t=201345
const testMinMacVlanVersion = "7.12.1"

func TestAccInterfaceMacVlanTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testMinMacVlanVersion) {
		t.Skipf("Test skipped, the minimum required version is %v", testMinMacVlanVersion)
		return
	}
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
							resource.TestCheckResourceAttr(testInterfaceMacVlanAddress, "disabled", "false"),
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
