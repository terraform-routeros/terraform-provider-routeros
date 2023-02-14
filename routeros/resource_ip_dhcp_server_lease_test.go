package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpDhcpServerLease = "routeros_ip_dhcp_server_lease.test_dhcp_lease"

func TestAccIpDhcpServerLeaseTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server/lease", "routeros_i[_dhcp_server_lease"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerLeaseConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpDhcpServerLeaseExists(testIpDhcpServerLease),
							resource.TestCheckResourceAttr(testIpDhcpServerLease, "address", "192.168.88.33"),
							resource.TestCheckResourceAttr(testIpDhcpServerLease, "mac_address", "AA:BB:CC:DD:EE:FF"),
							resource.TestCheckResourceAttr(testIpDhcpServerLease, "block_access", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDhcpServerLeaseExists(name string) resource.TestCheckFunc {
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

func testAccIpDhcpServerLeaseConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_dhcp_server_lease" "test_dhcp_lease" {
	address 	 = "192.168.88.33"
	mac_address	 = "AA:BB:CC:DD:EE:FF"
	block_access = true
  }

`
}
