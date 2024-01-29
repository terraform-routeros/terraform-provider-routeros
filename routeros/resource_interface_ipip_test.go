package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPIPAddress = "routeros_interface_ipip.ipip900"
const testIPIPName = "IPIP_900_TEST"

func TestAccInterfaceIPIPTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ipip", "routeros_interface_ipip"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIPIPConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPIPAddress),
							resource.TestCheckResourceAttr(testIPIPAddress, "name", testIPIPName),
						),
					},
				},
			})

		})
	}

}

func testAccInterfaceIPIPConfig() string {
	return providerConfig + `

resource "routeros_interface_ipip" "ipip900" {
	name      = "IPIP_900_TEST"
	remote_address = "127.0.0.1"
	disabled  = true
}
`
}
