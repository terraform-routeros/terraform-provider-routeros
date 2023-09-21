package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckLoggingExists(testSystemSimpleLoggingTask),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "action", "echo"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "disabled", "false"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "invalid", "false"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "prefix", "simple_prefix"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "topics.0", "snmp"),
							resource.TestCheckResourceAttr(testSystemSimpleLoggingTask, "topics.1", "gsm"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckLoggingExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccSystemLoggingConfig() string {
	return providerConfig + `
resource "routeros_system_logging" "simple_logging" {
    action = "echo"
    prefix = "simple_prefix"
    topics = ["snmp", "gsm"]
}
`
}
