package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceVxlanVteps = "routeros_interface_vxlan_vteps.test"

func TestAccInterfaceVxlanVtepsTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vxlan/vteps", "routeros_interface_vxlan_vteps"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVxlanVtepsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceVxlanVteps),
							resource.TestCheckResourceAttr(testInterfaceVxlanVteps, "interface", "vxlan2-test"),
							resource.TestCheckResourceAttr(testInterfaceVxlanVteps, "remote_ip", "192.168.10.10"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVxlanVtepsConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_interface_vxlan" "test-2" {
  name = "vxlan2-test"
  vni  = 11
}

resource "routeros_interface_vxlan_vteps" "test" {
  interface = routeros_interface_vxlan.test-2.name
  remote_ip = "192.168.10.10"
}
`, providerConfig)
}
