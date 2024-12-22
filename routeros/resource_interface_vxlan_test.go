package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceVxlan = "routeros_interface_vxlan.test"

func TestAccInterfaceVxlanTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vxlan", "routeros_interface_vxlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVxlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceVxlan),
							resource.TestCheckResourceAttr(testInterfaceVxlan, "name", "vxlan1-test"),
							resource.TestCheckResourceAttr(testInterfaceVxlan, "vni", "10"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVxlanConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_vxlan" "test" {
  name = "vxlan1-test"
  vni  = 10
}
`, providerConfig)
}
