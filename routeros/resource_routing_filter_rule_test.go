package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterRule = "routeros_routing_filter_rule.test"

func TestAccRoutingFilterRuleTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/filter/rule", "routeros_routing_filter_rule"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterRuleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterRule),
							resource.TestCheckResourceAttr(testRoutingFilterRule, "chain", "testChain"),
							resource.TestCheckResourceAttr(testRoutingFilterRule, "rule", "if (dst in 192.168.1.0/24 && dst-len>24) {set distance +1; accept} else {set distance -1; accept}"),
							resource.TestCheckResourceAttr(testRoutingFilterRule, "comment", "comment"),
							resource.TestCheckResourceAttr(testRoutingFilterRule, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterRuleConfig() string {
	return providerConfig + `
resource "routeros_routing_filter_rule" "test" {
  chain    = "testChain"
  rule     = "if (dst in 192.168.1.0/24 && dst-len>24) {set distance +1; accept} else {set distance -1; accept}"
  comment  = "comment"
  disabled = true
}
`
}
