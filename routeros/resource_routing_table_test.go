package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckRoutingTableExists(testRoutingTableAddress),
							resource.TestCheckResourceAttr(testRoutingTableAddress, "name", "to_ISP1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingTableExists(name string) resource.TestCheckFunc {
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

func testAccRoutingTableConfig() string {
	return providerConfig + `

resource "routeros_routing_table" "test_table" {
	name        = "to_ISP1"
	fib			= false
}
`
}
