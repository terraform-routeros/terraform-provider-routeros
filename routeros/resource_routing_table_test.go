package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingTableAddress = "routeros_routing_table.test_table"

func TestAccRoutingTableTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/table", "routeros_routing_table"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingTableConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingTableAddress),
							resource.TestCheckResourceAttr(testRoutingTableAddress, "name", "to_ISP1"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingTableConfig() string {
	return providerConfig + `

resource "routeros_routing_table" "test_table" {
	name        = "to_ISP1"
	fib			= false
}
`
}
