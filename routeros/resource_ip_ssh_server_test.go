package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSSHServerSettings = "routeros_ip_ssh_server.test"

func TestAccIpSSHServerSettingsTest_basic(t *testing.T) {
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
						Config: testAccIpSSHServerSettingsConfig(true, "both", 1024),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSSHServerSettings),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "allow_none_crypto", "true"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "always_allow_password_login", "true"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "forwarding_enabled", "both"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "host_key_size", "1024"),
						),
					},
					{
						Config: testAccIpSSHServerSettingsConfig(false, "no", 2048),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "allow_none_crypto", "false"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "always_allow_password_login", "false"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "forwarding_enabled", "no"),
							resource.TestCheckResourceAttr(testIpSSHServerSettings, "host_key_size", "2048"),
						),
					},
				},
			})
		})
	}
}

func testAccIpSSHServerSettingsConfig(b bool, fwd string, kSize int) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ssh_server" "test" {
	allow_none_crypto           = %v
	always_allow_password_login = %v
	forwarding_enabled          = "%v"
	host_key_size               = %v
}
`, providerConfig, b, b, fwd, kSize)
}
