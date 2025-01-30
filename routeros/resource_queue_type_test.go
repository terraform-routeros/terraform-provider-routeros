package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testQueueType = "routeros_queue_type.test"

func TestAccQueueTypeTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/queue/type", "routeros_queue_type"),
				Steps: []resource.TestStep{
					{
						Config: testAccQueueTypeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testQueueType),
							resource.TestCheckResourceAttr(testQueueType, "name", "pcq-test"),
							resource.TestCheckResourceAttr(testQueueType, "kind", "pcq"),
							resource.TestCheckResourceAttr(testQueueType, "pcq_rate", "0"),
							resource.TestCheckResourceAttr(testQueueType, "pcq_limit", "50"),
							resource.TestCheckResourceAttr(testQueueType, "pcq_classifier.0", "dst-address"),
						),
					},
				},
			})

		})
	}
}

func testAccQueueTypeConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_queue_type" "test" {
  name           = "pcq-test"
  kind           = "pcq"
  pcq_rate       = 0
  pcq_limit      = 50
  pcq_classifier = ["dst-address"]
}
`, providerConfig)
}
