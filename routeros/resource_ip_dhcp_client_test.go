package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpDhcpClientAddress = "routeros_ip_dhcp_client.test_dhcp"

func TestAccIpDhcpClientTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-client", "routeros_ip_dhcp_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpClientConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpDhcpClientExists(testIpDhcpClientAddress),
							resource.TestCheckResourceAttr(testIpDhcpClientAddress, "interface", "bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDhcpClientExists(name string) resource.TestCheckFunc {
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

func testAccIpDhcpClientConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_dhcp_client" "test_dhcp" {
	interface = "bridge"
  }

`
}
