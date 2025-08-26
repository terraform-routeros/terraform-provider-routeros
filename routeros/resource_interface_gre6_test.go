package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testGre6Address = "routeros_interface_gre6.gre_v6"

func TestAccInterfaceGre6Test_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/gre6", "routeros_interface_gre6"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceGre6Config(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testGre6Address),
							resource.TestCheckResourceAttr(testGre6Address, "name", "gre_v6"),
						),
					},
				},
			})

		})
	}

}

func testAccInterfaceGre6Config() string {
	return providerConfig + `

resource "routeros_interface_veth" "veth_v6" {
  name    = "veth_v6"
  address = ["2a02::1/64"]
}

resource "routeros_interface_gre6" "gre_v6" {
  name           = "gre_v6"
  remote_address = "2a02::2"
  disabled       = true
  depends_on  = [routeros_interface_veth.veth_v6]
}
`
}
