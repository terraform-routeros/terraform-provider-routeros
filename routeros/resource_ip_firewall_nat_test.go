package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckIPFirewallNatExists(testIPFirewallNat),
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

func testAccCheckIPFirewallNatExists(name string) resource.TestCheckFunc {
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
