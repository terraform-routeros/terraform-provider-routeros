package routeros

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeSettingsAddress = "routeros_interface_bridge_settings.test"

func TestAccInterfaceBridgeSettingsTest_basic(t *testing.T) {
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
						Config: testAccInterfaceBridgeSettingsConfig(name),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeSettingsAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeSettingsAddress, "id", "interface.bridge.settings"),
						),
					},
				},
			})
		})
	}
}

func testAccInterfaceBridgeSettingsConfig(testName string) string {
	if strings.Contains(testName, "API") {
		return providerConfig + `
resource "routeros_interface_bridge_settings" "test" {
	use_ip_firewall	= true
}
`
	}
	return providerConfig + `
resource "routeros_interface_bridge_settings" "test" {
	use_ip_firewall	= false
}
	`
}
