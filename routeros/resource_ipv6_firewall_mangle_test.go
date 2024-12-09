package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6FirewallMangle = "routeros_ipv6_firewall_mangle.data"

func TestAccIPv6FirewallMangleTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/firewall/mangle", "routeros_ipv6_firewall_mangle"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6FirewallMangleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6FirewallMangle),
							resource.TestCheckResourceAttr(testIPv6FirewallMangle, "chain", "prerouting"),
							resource.TestCheckResourceAttr(testIPv6FirewallMangle, "action", "mark-connection"),
							resource.TestCheckResourceAttr(testIPv6FirewallMangle, "new_connection_mark", "test-mark"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6FirewallMangleConfig() string {
	return providerConfig + `

resource "routeros_ipv6_firewall_mangle" "data" {
	chain = "prerouting"	
	action = "mark-connection"
	new_connection_mark = "test-mark"
	comment = "new-mangle-rule"
}

`
}
