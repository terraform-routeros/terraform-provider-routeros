package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolsBandwidthServer = "routeros_tool_bandwidth_server.test"

func TestAccToolsBandwidthServerTest_basic(t *testing.T) {
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
						Config: testAccToolsBandwidthServerConfig_none("none"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolsBandwidthServer),
							resource.TestCheckResourceAttr(testToolsBandwidthServer),
						),
					},
					{
						Config: testAccToolsBandwidthServerConfig_complex("all"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testToolsBandwidthServer),
						),
					},
				},
			})
		})
	}
}

func testAccToolsBandwidthServerConfig_none(acl string) string {
	return providerConfig + `

resource "routeros_tool_bandwidth_test_server" "test" {
	enabled = false
}
`
}

func testAccToolsBandwidthServerConfig_complex(acl string) string {
	return providerConfig + `

resource "routeros_tool_bandwidth_test_server" "test" {
	enabled = true
	authenticate = false
	max_sessions = 100
	allocate_udp_ports_from = 2000
}
`
}
