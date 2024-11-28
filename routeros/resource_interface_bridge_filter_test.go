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
							resource.TestCheckResourceAttr(testBridgeFilterRule, "action", "accept"),
						),
					},
				},
			})

		})
	}
}

func testAccBridgeFilterConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_filter" "rule" {
  chain             = "forward"
  action            = "accept"
  comment           = "test comment"
  log               = true
  log_prefix        = "log prefix here"
  disabled          = false
  dst_port          = "80"
  protocol          = "tcp"
  jump_target       = "test_target"
}
`
}
