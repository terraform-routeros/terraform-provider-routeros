package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpTrafficFlowIpfix = "routeros_ip_traffic_flow_ipfix.test"

func TestAccIpTrafficFlowIpfixTest_basic(t *testing.T) {
	t.Parallel()
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
						Config: testAccIpTrafficFlowIpfixConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpTrafficFlowIpfix),
							resource.TestCheckResourceAttr(testIpTrafficFlowIpfix, "nat_events", "true"),
							resource.TestCheckResourceAttr(testIpTrafficFlowIpfix, "dst_port", "false"),
						),
					},
					{
						Config: testAccIpTrafficFlowIpfixConfigRollback(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpTrafficFlowIpfix),
							resource.TestCheckResourceAttr(testIpTrafficFlowIpfix, "nat_events", "false"),
							resource.TestCheckResourceAttr(testIpTrafficFlowIpfix, "dst_port", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIpTrafficFlowIpfixConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_traffic_flow_ipfix" "test" {
  nat_events = true
  dst_port   = false
}
`, providerConfig)
}

func testAccIpTrafficFlowIpfixConfigRollback() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_traffic_flow_ipfix" "test" {
  nat_events = false
  dst_port   = true
}
`, providerConfig)
}
