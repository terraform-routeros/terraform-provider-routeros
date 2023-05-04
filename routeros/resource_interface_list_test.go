package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceListAddress = "routeros_interface_list.test_list"

func TestAccInterfaceListTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/list", "routeros_interface_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceListConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceListExists(testInterfaceListAddress),
							resource.TestCheckResourceAttr(testInterfaceListAddress, "name", "test_list"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceListExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceListConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_list" "test_list" {
	name      = "test_list"
}
`
}
