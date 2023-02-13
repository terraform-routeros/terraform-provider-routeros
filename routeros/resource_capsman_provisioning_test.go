package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManProvisioningAddress = "routeros_capsman_provisioning.test_provisioning"

func TestAccCapsManProvisioningTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/provisioning", "routeros_capsman_provisioning"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManProvisioningConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManProvisioningExists(testInterfaceBridgeAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddress, "name", "test_provisioning"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckCapsManProvisioningExists(name string) resource.TestCheckFunc {
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

func testAccCapsManProvisioningConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_provisioning" "test_provisioning" {
	action   = "create-disabled"
  }

`
}
