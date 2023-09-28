package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpServerAddress = "routeros_ip_dhcp_server.test_dhcp"

func TestAccIpDhcpServerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server", "routeros_ip_dhcp_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpServerAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerAddress, "interface", "bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpServerConfig() string {
	return providerConfig + `
resource "routeros_ip_dhcp_server" "test_dhcp" {
	name	     = "test_dhcp_server"
	interface    = "bridge"
	disabled     = true
	address_pool = "dhcp"
  }

`
}
