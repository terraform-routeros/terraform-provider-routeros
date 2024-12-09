package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6FirewallNat = "routeros_ipv6_firewall_nat.data"

func TestAccIPv6FirewallNatTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/firewall/nat", "routeros_ipv6_firewall_nat"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6FirewallNatConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6FirewallNat),
							resource.TestCheckResourceAttr(testIPv6FirewallNat, "chain", "srcnat"),
							resource.TestCheckResourceAttr(testIPv6FirewallNat, "action", "masquerade"),
							resource.TestCheckResourceAttr(testIPv6FirewallNat, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6FirewallNatConfig() string {
	return providerConfig + `

resource "routeros_ipv6_firewall_nat" "data" {
	chain = "srcnat"	
	action = "masquerade"
	disabled = true
	comment = "new-nat-rule"
}

`
}
