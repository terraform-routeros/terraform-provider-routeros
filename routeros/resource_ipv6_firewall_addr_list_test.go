package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckIPv6FirewallAddrListExists(testIPv6FirewallAddrList),
							resource.TestCheckResourceAttr(testIPv6FirewallAddrList, "list", "test-addr-list"),
							resource.TestCheckResourceAttr(testIPv6FirewallAddrList, "address", "123:dead:beaf::/64"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIPv6FirewallAddrListExists(name string) resource.TestCheckFunc {
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

func testAccIPv6FirewallAddrListConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ipv6_firewall_addr_list" "data" {
	list = "test-addr-list"
	address = "123:dead:beaf::/64"
	timeout = "5m"
}

`
}
