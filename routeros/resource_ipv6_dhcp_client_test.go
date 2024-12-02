package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6DhcpClient = "routeros_ipv6_dhcp_client.client"

func TestAccIPv6DhcpClient_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-client", "routeros_ipv6_dhcp_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6DhcpClientConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpClient),
							resource.TestCheckResourceAttr(testIPv6DhcpClient, "interface", "ether1"),
							resource.TestCheckResourceAttr(testIPv6DhcpClient, "pool_name", "inet-provider-pool"),
							resource.TestCheckResourceAttr(testIPv6DhcpClient, "request.0", "prefix"),
							resource.TestCheckResourceAttr(testIPv6DhcpClient, "prefix_hint", "::/60"),
						),
					},
				},
			})
		})
	}
}

func testAccIPv6DhcpClientConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_client" "client" {	
  request            = ["prefix"]
  pool_name          = "inet-provider-pool"
  pool_prefix_length = 64
  interface          = "ether1"
  prefix_hint        = "::/60"
}

`
}
