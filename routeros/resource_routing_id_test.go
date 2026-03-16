package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingId = "routeros_routing_id.test"

func TestAccRoutingIdTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/id", "routeros_routing_id"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingIdConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingId),
							resource.TestCheckResourceAttr(testRoutingId, "name", "router-id-test"),
							resource.TestCheckResourceAttr(testRoutingId, "router_id", "10.10.10.10"),
							resource.TestCheckResourceAttr(testRoutingId, "select_dynamic_id", "any"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIdConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_id" "test" {
  name              = "router-id-test"
  router_id         = "10.10.10.10"
  select_dynamic_id = "any"
}
`, providerConfig)
}
