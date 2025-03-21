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
						Config: testAccIpSettingsConfig("l4"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSettings),
							resource.TestCheckResourceAttr(testIpSettings, "ipv4_multipath_hash_policy", "l4"),
						),
					},
					{
						Config: testAccIpSettingsConfig("l3-inner"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSettings),
							resource.TestCheckResourceAttr(testIpSettings, "ipv4_multipath_hash_policy", "l3-inner"),
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
  ipv4_multipath_hash_policy = "%v"
}
`, providerConfig, param)
}
