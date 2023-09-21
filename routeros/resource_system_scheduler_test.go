package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemSchedulerTask = "routeros_scheduler.test_task"

func TestAccSystemSchedulerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/scheduler", "routeros_scheduler"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemSchedulerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemSchedulerTask),
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

func testAccSystemSchedulerConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_scheduler" "test_task" {
	name = "TestTask"
	on_event="script1"
	disabled = true
    policy = ["ftp", "read", "write"]
}
`
}
