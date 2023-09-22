package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPFirewallNat = "routeros_firewall_nat.data"

func TestAccIPFirewallNatTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/nat", "routeros_firewall_nat"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallNatConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallNat),
							resource.TestCheckResourceAttr(testIPFirewallNat, "chain", "srcnat"),
							resource.TestCheckResourceAttr(testIPFirewallNat, "action", "masquerade"),
							resource.TestCheckResourceAttr(testIPFirewallNat, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIPFirewallNatConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_firewall_nat" "data" {
	chain = "srcnat"	
	action = "masquerade"
	disabled = true
	comment = "new-nat-rule"
}

`
}
