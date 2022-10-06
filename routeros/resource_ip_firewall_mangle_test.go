package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIPFirewallMangle = "routeros_firewall_mangle.data"

func TestAccIPFirewallMangleTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/mangle", "routeros_firewall_mangle"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallMangleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIPFirewallMangleExists(testIPFirewallMangle),
							resource.TestCheckResourceAttr(testIPFirewallMangle, "chain", "prerouting"),
							resource.TestCheckResourceAttr(testIPFirewallMangle, "action", "mark-connection"),
							resource.TestCheckResourceAttr(testIPFirewallMangle, "new_connection_mark", "test-mark"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIPFirewallMangleExists(name string) resource.TestCheckFunc {
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

func testAccIPFirewallMangleConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_firewall_mangle" "data" {
	chain = "prerouting"	
	action = "mark-connection"
	new_connection_mark = "test-mark"
	comment = "new-mangle-rule"
}

`
}
