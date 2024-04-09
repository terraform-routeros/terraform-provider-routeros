package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUPNPSettings = "routeros_ip_upnp.test"

func TestAccUPNPSettingsTest_basic(t *testing.T) {
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
						Config: testAccUPNPSettingsConfig(true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUPNPSettings),
							resource.TestCheckResourceAttr(testUPNPSettings, "allow_disable_external_interface", "true"),
							resource.TestCheckResourceAttr(testUPNPSettings, "enabled", "true"),
							resource.TestCheckResourceAttr(testUPNPSettings, "show_dummy_rule", "true"),
						),
					},
					{
						Config: testAccUPNPSettingsConfig(false),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testUPNPSettings, "allow_disable_external_interface", "false"),
							resource.TestCheckResourceAttr(testUPNPSettings, "enabled", "false"),
							resource.TestCheckResourceAttr(testUPNPSettings, "show_dummy_rule", "false"),
						),
					},
				},
			})
		})
	}
}

func testAccUPNPSettingsConfig(b bool) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_upnp" "test" {
	allow_disable_external_interface = %v
	enabled                          = %v
	show_dummy_rule                  = %v
}
`, providerConfig, b, b, b)
}
