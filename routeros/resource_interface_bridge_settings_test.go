package routeros

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckInterfaceBridgeSettingsExists(testInterfaceBridgeSettingsAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeSettingsAddress, "id", "interface.bridge.settings"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckInterfaceBridgeSettingsExists(name string) resource.TestCheckFunc {
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
