package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckIpDhcpServerExists(testIpDhcpServerAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerAddress, "interface", "bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDhcpServerExists(name string) resource.TestCheckFunc {
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

func testAccIpDhcpServerConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_dhcp_server" "test_dhcp" {
	name	     = "test_dhcp_server"
	interface    = "bridge"
	disabled     = true
	address_pool = "dhcp"
  }

`
}
