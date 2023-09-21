package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpServerNetworkAddress = "routeros_ip_dhcp_server_network.test_dhcp"

func TestAccIpDhcpServerNetworkTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server/network", "routeros_ip_dhcp_server_network"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerNetworkConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpServerNetworkAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerNetworkAddress, "address", "192.168.1.0/24"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpServerNetworkConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_dhcp_server_network" "test_dhcp" {
	address    = "192.168.1.0/24"
  }

`
}
