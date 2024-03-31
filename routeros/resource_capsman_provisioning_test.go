package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testCapsManProvisioningMaxVersion = "7.12.2"
const testCapsManProvisioningAddress = "routeros_capsman_provisioning.test_provisioning"

func TestAccCapsManProvisioningTest_basic(t *testing.T) {
	if !testCheckMaxVersion(t, testCapsManProvisioningMaxVersion) {
		t.Logf("Test skipped, the maximum supported version is %v", testCapsManProvisioningMaxVersion)
		return
	}

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
							testResourcePrimaryInstanceId(testCapsManProvisioningAddress),
							resource.TestCheckResourceAttr(testCapsManProvisioningAddress, "name_prefix", "cap-"),
						),
					},
				},
			})

		})
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
