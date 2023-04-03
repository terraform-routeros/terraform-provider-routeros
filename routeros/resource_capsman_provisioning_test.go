package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckCapsManProvisioningExists(testCapsManProvisioningAddress),
							resource.TestCheckResourceAttr(testCapsManProvisioningAddress, "name_prefix", "cap-"),
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
	return providerConfig + `

resource "routeros_capsman_configuration" "test_configuration" {
	name = "cfg1"
}

resource "routeros_capsman_provisioning" "test_provisioning" {
	master_configuration = "cfg1"
	action               = "create-disabled"
	name_prefix          = "cap-"

	depends_on = [
		routeros_capsman_configuration.test_configuration,
	]
}

`
}
