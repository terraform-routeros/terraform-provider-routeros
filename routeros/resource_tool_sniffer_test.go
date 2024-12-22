package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolSniffer = "routeros_tool_sniffer.test"

func TestAccToolSnifferTest_basic(t *testing.T) {
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
						Config: testAccToolSnifferConfig("or"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolSniffer),
							resource.TestCheckResourceAttr(testToolSniffer, "streaming_enabled", "true"),
							resource.TestCheckResourceAttr(testToolSniffer, "streaming_server", "192.168.88.5:37008"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_stream", "true"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_interface.0", "ether3"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_direction", "rx"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_operator_between_entries", "or"),
						),
					},
					{
						Config: testAccToolSnifferConfig("and"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolSniffer),
							resource.TestCheckResourceAttr(testToolSniffer, "streaming_enabled", "true"),
							resource.TestCheckResourceAttr(testToolSniffer, "streaming_server", "192.168.88.5:37008"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_stream", "true"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_interface.0", "ether3"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_direction", "rx"),
							resource.TestCheckResourceAttr(testToolSniffer, "filter_operator_between_entries", "and"),
						),
					},
				},
			})

		})
	}
}

func testAccToolSnifferConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_sniffer" "test" {
  streaming_enabled = true
  streaming_server  = "192.168.88.5:37008"
  filter_stream     = true

  filter_interface                = ["ether3"]
  filter_direction                = "rx"
  filter_operator_between_entries = "%v"
}
`, providerConfig, param)
}
