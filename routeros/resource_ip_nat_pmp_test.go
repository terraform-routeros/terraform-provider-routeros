package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testNatPmpMinVersion = "7.13"
const testNatPmpSettings = "routeros_ip_nat_pmp.test"

func TestAccNatPmpSettingsTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testNatPmpMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testNatPmpMinVersion)
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
				Steps: []resource.TestStep{
					{
						Config: testAccNatPmpSettingsConfig(true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testNatPmpSettings),
							resource.TestCheckResourceAttr(testNatPmpSettings, "enabled", "true"),
						),
					},
					{
						Config: testAccNatPmpSettingsConfig(false),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testNatPmpSettings, "enabled", "false"),
						),
					},
				},
			})
		})
	}
}

func testAccNatPmpSettingsConfig(b bool) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_nat_pmp" "test" {
	enabled = %v
}
`, providerConfig, b)
}
