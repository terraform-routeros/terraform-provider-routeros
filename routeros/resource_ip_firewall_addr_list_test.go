package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPFirewallrAddrList = "routeros_firewall_addr_list.data"

func TestAccIPFirewallAddrListTest_basic(t *testing.T) {
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
							testResourcePrimaryInstanceId(testIPFirewallrAddrList),
							resource.TestCheckResourceAttr(testIPFirewallrAddrList, "list", "test-addr-list"),
							resource.TestCheckResourceAttr(testIPFirewallrAddrList, "address", "192.168.0.0/23"),
						),
					},
				},
			})

		})
	}
}

func testAccIPFirewallAddrListConfig() string {
	return providerConfig + `

resource "routeros_firewall_addr_list" "data" {
	list = "test-addr-list"	
	address = "192.168.0.0-192.168.1.255"
	timeout = "5m"
}

`
}
