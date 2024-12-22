package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecSettings = "routeros_ip_ipsec_settings.test"

func TestAccIpIpsecSettingsTest_basic(t *testing.T) {
	// t.Parallel()
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
						Config: testAccIpIpsecSettingsConfig("true", "10s"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecSettings),
							resource.TestCheckResourceAttr(testIpIpsecSettings, "xauth_use_radius", "true"),
							resource.TestCheckResourceAttr(testIpIpsecSettings, "interim_update", "10s"),
						),
					},
					{
						Config: testAccIpIpsecSettingsConfig("false", "0s"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecSettings),
							resource.TestCheckResourceAttr(testIpIpsecSettings, "xauth_use_radius", "false"),
							resource.TestCheckResourceAttr(testIpIpsecSettings, "interim_update", "0s"),
						),
					},
				},
			})

		})
	}
}

func testAccIpIpsecSettingsConfig(param1, param2 string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_settings" "test" {
  xauth_use_radius = %v
  interim_update   = "%v"
}
`, providerConfig, param1, param2)
}
