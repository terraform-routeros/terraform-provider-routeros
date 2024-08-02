package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemSimpleLoggingTask = "routeros_system_logging.simple_logging"

func TestAccSystemLoggingTest_basic(t *testing.T) {
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
						Config: testAccSystemLoggingConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemSimpleLoggingTask),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "action", "echo"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "disabled", "false"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "invalid", "false"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "prefix", "simple_prefix"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "topics.0", "gsm"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "topics.1", "snmp"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemLoggingConfig() string {
	return providerConfig + `
resource "routeros_system_logging" "simple_logging" {
    action = "echo"
    prefix = "simple_prefix"
    topics = ["snmp", "gsm"]
}

resource "routeros_system_logging" "info" {
	action = "echo"
	prefix = ":Info"
	topics = ["info", "!wireguard"]
}
`
}
