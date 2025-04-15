package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testQueueTree = "routeros_queue_tree.test"

func TestAccQueueTreeTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/queue/tree", "routeros_queue_tree"),
				Steps: []resource.TestStep{
					{
						Config: testAccQueueTreeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testQueueTree),
							resource.TestCheckResourceAttr(testQueueTree, "name", "qt-test"),
							resource.TestCheckResourceAttr(testQueueTree, "parent", "global"),
							resource.TestCheckResourceAttr(testQueueTree, "packet_mark.0", "pmark-test"),
							resource.TestCheckResourceAttr(testQueueTree, "max_limit", "10000000"),
						),
					},
				},
			})

		})
	}
}

func testAccQueueTreeConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_queue_tree" "test" {
  name        = "qt-test"
  parent      = "global"
  max_limit   = "10M"
  packet_mark = ["pmark-test"]
}
`, providerConfig)
}
