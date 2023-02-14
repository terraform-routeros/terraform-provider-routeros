package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManConfigurationAddress = "routeros_capsman_configuration.test_configuration"

func TestAccCapsManConfigurationTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/configuration", "routeros_capsman_configuration"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManConfigurationConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManConfigurationExists(testCapsManConfigurationAddress),
							resource.TestCheckResourceAttr(testCapsManConfigurationAddress, "name", "test_configuration"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManConfigurationExists(name string) resource.TestCheckFunc {
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

func testAccCapsManConfigurationConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_configuration" "test_configuration" {
	name   = "test_configuration"
  }

`
}
