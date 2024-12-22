package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRule = "routeros_routing_rule.test"

func TestAccRoutingRuleTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRuleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRule),
							resource.TestCheckResourceAttr(testRoutingRule, "dst_address", "192.168.1.0/24"),
							resource.TestCheckResourceAttr(testRoutingRule, "action", "lookup-only-in-table"),
							resource.TestCheckResourceAttr(testRoutingRule, "interface", "ether1"),
						),
					},
					{
						Config: testAccRoutingRuleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRule),
							resource.TestCheckResourceAttr(testRoutingRule, "dst_address", "192.168.1.0/24"),
							resource.TestCheckResourceAttr(testRoutingRule, "action", "lookup-only-in-table"),
							resource.TestCheckResourceAttr(testRoutingRule, "interface", "ether1"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRuleConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_routing_rule" "test" {
  dst_address = "192.168.1.0/24"
  action      = "lookup-only-in-table"
  interface   = "ether1"
}
`, providerConfig)
}
