package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolNetwatch = "routeros_tool_netwatch.test"

func TestAccToolNetwatchTest_basic(t *testing.T) {
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
						Config: testAccToolNetwatchConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "name", "watch-google-pdns"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "8.8.8.8"),
							resource.TestCheckResourceAttr(testToolNetwatch, "interval", "30s"),
							resource.TestCheckResourceAttr(testToolNetwatch, "up_script", ":log info \"Ping to 8.8.8.8 successful\""),
						),
					},
					{
						Config: testAccToolNetwatchConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "name", "watch-google-pdns"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "8.8.8.8"),
							resource.TestCheckResourceAttr(testToolNetwatch, "interval", "30s"),
							resource.TestCheckResourceAttr(testToolNetwatch, "up_script", ":log info \"Ping to 8.8.8.8 successful\""),
						),
					},
				},
			})

		})
	}
}

func testAccToolNetwatchConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_tool_netwatch" "test" {
  name      = "watch-google-pdns"
  type      = "icmp"
  host      = "8.8.8.8"
  interval  = "30s"
  up_script = ":log info \"Ping to 8.8.8.8 successful\""
  thr_max   = "400ms"
}
`, providerConfig)
}
