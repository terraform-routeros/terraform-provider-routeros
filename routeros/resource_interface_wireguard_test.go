package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWireguardAddress = "routeros_interface_wireguard.test_wg_interface"

func TestAccInterfaceWireguardTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/wireguard", "routeros_interface_wireguard"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWireguardConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWireguardAddress),
							resource.TestCheckResourceAttr(testInterfaceWireguardAddress, "listen_port", "13231"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWireguardConfig() string {
	return providerConfig + `

resource "routeros_interface_wireguard" "test_wg_interface" {
	name   		= "test_wg_interface"
	listen_port = "13231"
  }

`
}
