package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpDhcpServerNetworkAddress = "routeros_dhcp_server_network.test_dhcp"

func TestAccIpDhcpServerNetworkTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/ip/dhcp-server/network", "routeros_dhcp_server_network"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerNetworkConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpDhcpServerNetworkExists(testIpDhcpServerNetworkAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerNetworkAddress, "address", "192.168.1.0/24"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDhcpServerNetworkExists(name string) resource.TestCheckFunc {
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

func testAccIpDhcpServerNetworkConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_dhcp_server_network" "test_dhcp" {
	address    = "192.168.1.0/24"
  }

`
}
