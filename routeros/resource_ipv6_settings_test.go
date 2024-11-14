package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6Settings = "routeros_ipv6_settings.settings"

func TestAccIpv6SettingsTest_basic(t *testing.T) {
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
						Config: testAccIpv6SettingsConfig("no"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6Settings),
							resource.TestCheckResourceAttr(testIpv6Settings, "accept_redirects", "no"),
						),
					},
					{
						Config: testAccIpv6SettingsConfig("yes-if-forwarding-disabled"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6Settings),
							resource.TestCheckResourceAttr(testIpv6Settings, "accept_redirects", "yes-if-forwarding-disabled"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6SettingsConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_settings" "settings" {
  accept_redirects = "%v"
}
`, providerConfig, param)
}
