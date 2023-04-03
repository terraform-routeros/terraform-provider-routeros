package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManSecurityAddress = "routeros_capsman_security.test_security"

func TestAccCapsManSecurityTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/security", "routeros_capsman_security"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManSecurityConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManSecurityExists(testCapsManSecurityAddress),
							resource.TestCheckResourceAttr(testCapsManSecurityAddress, "name", "test_security"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckCapsManSecurityExists(name string) resource.TestCheckFunc {
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

func testAccCapsManSecurityConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_security" "test_security" {
	name   = "test_security"
  }

`
}
