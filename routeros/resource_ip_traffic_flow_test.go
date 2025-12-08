package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpTrafficFlow = "routeros_ip_traffic_flow.test"

func TestAccIpTrafficFlowTest_basic(t *testing.T) {
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
						Config: testAccIpTrafficFlowConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpTrafficFlow),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "packet_sampling", "true"),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "sampling_interval", "2222"),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "sampling_space", "1111"),
						),
					},
					{
						Config: testAccIpTrafficFlowConfigRollback(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpTrafficFlow),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "packet_sampling", "false"),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "sampling_interval", "0"),
							resource.TestCheckResourceAttr(testIpTrafficFlow, "sampling_space", "0"),
						),
					},
				},
			})

		})
	}
}

func testAccIpTrafficFlowConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_traffic_flow" "test" {
  packet_sampling   = true
  sampling_interval = 2222
  sampling_space    = 1111
}
`, providerConfig)
}

func testAccIpTrafficFlowConfigRollback() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_traffic_flow" "test" {
  packet_sampling   = false
  sampling_interval = 0
  sampling_space    = 0
}
`, providerConfig)
}
