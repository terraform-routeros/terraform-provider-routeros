package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6FirewallAddrList = "routeros_ipv6_firewall_addr_list.data"

func TestAccIPv6FirewallAddrListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/firewall/address-list", "routeros_ipv6_firewall_addr_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6FirewallAddrListConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6FirewallAddrList),
							resource.TestCheckResourceAttr(testIPv6FirewallAddrList, "list", "test-addr-list"),
							resource.TestCheckResourceAttr(testIPv6FirewallAddrList, "address", "123:dead:beaf::/64"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6FirewallAddrListConfig() string {
	return providerConfig + `

resource "routeros_ipv6_firewall_addr_list" "data" {
	list = "test-addr-list"
	address = "123:dead:beaf::/64"
	timeout = "5m"
}

`
}
