package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManManagerAddress = "routeros_capsman_manager.test_manager"

func TestAccCapsManManagerTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCapsManManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCapsManManagerExists(testCapsManManagerAddress),
					resource.TestCheckResourceAttr(testCapsManManagerAddress, "enabled", "true"),
				),
			},
		},
	})
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
