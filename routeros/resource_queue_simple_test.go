package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testQueueSimple = "routeros_queue_simple.test"

func TestAccQueueSimpleTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/queue/simple", "routeros_queue_simple"),
				Steps: []resource.TestStep{
					{
						Config: testAccQueueSimpleConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testQueueSimple),
							resource.TestCheckResourceAttr(testQueueSimple, "name", "server"),
							resource.TestCheckResourceAttr(testQueueSimple, "target.0", "10.1.1.1/32"),
							resource.TestCheckResourceAttr(testQueueSimple, "max_limit", "0/0"),
						),
					},
				},
			})

		})
	}
}

func testAccQueueSimpleConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_queue_simple" "test" {
  name      = "server"
  target    = ["10.1.1.1/32"]
  max_limit = "0/0"
}
`, providerConfig)
}
