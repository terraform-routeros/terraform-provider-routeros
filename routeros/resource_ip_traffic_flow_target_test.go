package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpTrafficFlowTarget = "routeros_ip_traffic_flow_target.test"

func TestAccIpTrafficFlowTargetTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/traffic-flow/target", "routeros_ip_traffic_flow_target"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpTrafficFlowTargetConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpTrafficFlowTarget),
							resource.TestCheckResourceAttr(testIpTrafficFlowTarget, "dst_address", "192.168.0.2"),
							resource.TestCheckResourceAttr(testIpTrafficFlowTarget, "port", "2055"),
							resource.TestCheckResourceAttr(testIpTrafficFlowTarget, "version", "9"),
						),
					},
				},
			})

		})
	}
}

func testAccIpTrafficFlowTargetConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_traffic_flow_target" "test" {
  dst_address = "192.168.0.2"
  port        = 2055
  version     = "9"
}
`, providerConfig)
}
