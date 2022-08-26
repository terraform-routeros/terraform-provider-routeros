package routeros

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testSystemSchedulerTask = "routeros_scheduler.test_task"

func TestAccSystemSchedulerTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/system/scheduler", "routeros_scheduler"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemSchedulerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckSystemSchedulerExists(testSystemSchedulerTask),
							resource.TestCheckResourceAttr(testSystemSchedulerTask, "disabled", "true"),
							resource.TestCheckResourceAttr(testSystemSchedulerTask, "name", "TestTask"),
							resource.TestCheckResourceAttr(testSystemSchedulerTask, "on_event", "script1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckSystemSchedulerExists(name string) resource.TestCheckFunc {
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

func testAccSystemSchedulerConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_scheduler" "test_task" {
	name = "TestTask"
	on_event="script1"
	disabled = true
}
`
}
