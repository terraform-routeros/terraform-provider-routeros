package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6DhcpServer = "routeros_ipv6_dhcp_server.test"

func TestAccIpv6DhcpServerTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-server", "routeros_ipv6_dhcp_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpv6DhcpServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6DhcpServer),
							resource.TestCheckResourceAttr(testIpv6DhcpServer, "address_pool", "test-pool-0"),
							resource.TestCheckResourceAttr(testIpv6DhcpServer, "interface", "bridge"),
							resource.TestCheckResourceAttr(testIpv6DhcpServer, "lease_time", "1m"),
							resource.TestCheckResourceAttr(testIpv6DhcpServer, "name", "test-dhcpv6"),
							resource.TestCheckResourceAttr(testIpv6DhcpServer, "preference", "128"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6DhcpServerConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_pool" "pool-0" {
  name          = "test-pool-0"
  prefix        = "2001:db8:40::/48"
  prefix_length = 64
}

resource "routeros_ipv6_dhcp_server" "test" {
  address_pool = routeros_ipv6_pool.pool-0.name
  interface    = "bridge"
  lease_time   = "1m"
  name         = "test-dhcpv6"
  preference   = 128
}
`, providerConfig)
}
