package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManManagerAddress = "routeros_capsman_manager.test_manager"

func TestAccCapsManManagerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/manager", "routeros_capsman_manager"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManManagerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManManagerExists(testCapsManManagerAddress),
							resource.TestCheckResourceAttr(testCapsManManagerAddress, "name", "test_manager"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManManagerExists(address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[address]
		if !ok {
			return fmt.Errorf("not found: %s", address)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccCapsManManagerConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_manager" "test_manager" {
	enabled   = "true"
  }

`
}
