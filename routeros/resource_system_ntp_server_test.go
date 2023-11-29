package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemNtpServer = "routeros_system_ntp_server.test"

func TestAccSystemNtpServerTest_basic(t *testing.T) {
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
						Config: testAccSystemNtpServerConfig("true", "10"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpServer),
							resource.TestCheckResourceAttr(testSystemNtpServer, "enabled", "true"),
							resource.TestCheckResourceAttr(testSystemNtpServer, "local_clock_stratum", "10"),
						),
					},
					{
						Config: testAccSystemNtpServerConfig("false", "3"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemNtpServer, "enabled", "false"),
							resource.TestCheckResourceAttr(testSystemNtpServer, "local_clock_stratum", "3"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemNtpServerConfig(param, stratum string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_ntp_server" "test" {
	enabled             = %v
	broadcast           = %v
	multicast           = %v
	manycast            = %v
	use_local_clock     = %v
	local_clock_stratum = %v
}
`, providerConfig, param, param, param, param, param, stratum)
}
