package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWirelessSecurityProfiles = "routeros_interface_wireless_security_profiles.test"

func TestAccInterfaceWirelessSecurityProfilesTest_basic(t *testing.T) {
	if !testCheckMaxVersion(t, testCapsManChannelMaxVersion) {
		t.Logf("Test skipped, the maximum supported version is %v", testCapsManChannelMaxVersion)
		return
	}

	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/wireless/security-profiles",
					"routeros_interface_wireless_security_profiles"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWirelessSecurityProfilesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWirelessSecurityProfiles),
							resource.TestCheckResourceAttr(testInterfaceWirelessSecurityProfiles, "name", "test-profile"),
							resource.TestCheckResourceAttr(testInterfaceWirelessSecurityProfiles, "mode", "dynamic-keys"),
							resource.TestCheckResourceAttr(testInterfaceWirelessSecurityProfiles, "wpa_pre_shared_key", "wpa_psk_key"),
							resource.TestCheckResourceAttr(testInterfaceWirelessSecurityProfiles, "wpa2_pre_shared_key", "wpa2_psk_key"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWirelessSecurityProfilesConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_wireless_security_profiles" "test" {
  name                 = "test-profile"
  mode                 = "dynamic-keys"
  authentication_types = ["wpa-psk", "wpa2-psk"]
  wpa_pre_shared_key   = "wpa_psk_key"
  wpa2_pre_shared_key  = "wpa2_psk_key"
}
`, providerConfig)
}
