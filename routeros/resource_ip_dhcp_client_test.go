package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testIpDhcpClientAddress),
							resource.TestCheckResourceAttr(testIpDhcpClientAddress, "interface", "bridge"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpClientConfig() string {
	return providerConfig + `

resource "routeros_ip_dhcp_client" "test_dhcp" {
	interface = "bridge"
  }

`
}
