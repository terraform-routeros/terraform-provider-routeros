package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIPFirewallrAddrList = "routeros_firewall_addr_list.data"

func TestAccIPFirewallAddrListTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/address-list", "routeros_firewall_addr_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallAddrListConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIPFirewallAddrListExists(testIPFirewallrAddrList),
							resource.TestCheckResourceAttr(testIPFirewallrAddrList, "list", "test-addr-list"),
							resource.TestCheckResourceAttr(testIPFirewallrAddrList, "address", "192.168.0.0/23"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIPFirewallAddrListExists(name string) resource.TestCheckFunc {
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

func testAccIPFirewallAddrListConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_firewall_addr_list" "data" {
	list = "test-addr-list"	
	address = "192.168.0.0-192.168.1.255"
	timeout = "5m"
}

`
}
