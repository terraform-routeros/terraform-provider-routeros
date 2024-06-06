package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemLoggingAction = "routeros_system_logging_action.action"

func TestAccSystemLoggingActionTest_basic(t *testing.T) {
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
						Config: testAccSystemLoggingActionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemLoggingAction),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "name", "action1"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "default", "false"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "target", "remote"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "remote", "192.168.1.1"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "bsd_syslog", "true"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_facility", "user"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_severity", "notice"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_time_format", "iso8601"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemLoggingActionConfig() string {
	return providerConfig + `
resource "routeros_system_logging_action" "action" {
	name               = "action1"
    target             = "remote"
	remote             = "192.168.1.1"
    bsd_syslog         = true
	syslog_facility    = "user"
	syslog_severity    = "notice"
	syslog_time_format = "iso8601"
}
`
}
