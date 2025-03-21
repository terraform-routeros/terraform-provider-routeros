package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSettings = "routeros_ip_settings.settings"

func TestAccIpSettingsTest_basic(t *testing.T) {
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
						Config: testAccIpSettingsConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSettings),
							resource.TestCheckResourceAttr(testIpSettings, "allow_fast_path", "false"),
						),
					},
					{
						Config: testAccIpSettingsConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSettings),
							resource.TestCheckResourceAttr(testIpSettings, "allow_fast_path", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSettingsConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_settings" "settings" {
  allow_fast_path = %v
}
`, providerConfig, param)
}
