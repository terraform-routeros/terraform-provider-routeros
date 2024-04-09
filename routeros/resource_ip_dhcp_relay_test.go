package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpRelayAddress = "routeros_ip_dhcp_relay.test"

func TestAccIpDhcpRelayTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-relay", "routeros_ip_dhcp_relay"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpRelayConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpRelayAddress),
							resource.TestCheckResourceAttr(testIpDhcpRelayAddress, "name", "test relay"),
							resource.TestCheckResourceAttr(testIpDhcpRelayAddress, "interface", "ether1"),
							resource.TestCheckResourceAttr(testIpDhcpRelayAddress, "dhcp_server", "0.0.0.1"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpRelayConfig() string {
	return providerConfig + `

resource "routeros_ip_dhcp_relay" "test" {
	name        = "test relay"
	interface   = "ether1"
	dhcp_server = "0.0.0.1"
  }

`
}
