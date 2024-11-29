package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testBridgeFilterRule = "routeros_interface_bridge_filter.rule"

func TestAccBridgeFilterTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bridge/filter", "routeros_interface_bridge_filter"),

				Steps: []resource.TestStep{
					{
						Config: testAccBridgeFilterConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testBridgeFilterRule),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "chain", "forward"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "action", "drop"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "log_prefix", "Blocking MS broadcast"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "comment", "HSIA - Block MS broadcast"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "ip_protocol", "udp"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "dst_port", "135-137"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "mac_protocol", "ip"),
							resource.TestCheckResourceAttr(testBridgeFilterRule, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccBridgeFilterConfig() string {
	return providerConfig + `
resource "routeros_interface_bridge_filter" "rule" {
  chain        = "forward"
  action       = "drop"
  log_prefix   = "Blocking MS broadcast"
  comment      = "HSIA - Block MS broadcast"
  ip_protocol  = "udp"
  dst_port     = "135-137"
  mac_protocol = "ip"
  disabled     = true
}
`
}
