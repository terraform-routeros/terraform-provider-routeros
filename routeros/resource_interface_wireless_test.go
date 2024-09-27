package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWireless = "routeros_interface_wireless.test"

func TestAccInterfaceWirelessTest_basic(t *testing.T) {
	t.Logf("A device with WiFi interface is required for the test")
	return

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/wireless", "routeros_interface_wireless"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWirelessConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWireless),
							resource.TestCheckResourceAttr(testInterfaceWireless, "", ""),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWirelessConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_wireless_security_profiles" "test" {
  name                 = "test-profile"
  mode                 = "dynamic-keys"
  authentication_types = ["wpa-psk", "wpa2-psk"]
  wpa_pre_shared_key   = "wpa_psk_key"
  wpa2_pre_shared_key  = "wpa2_psk_key"
}

resource "routeros_interface_wireless" "test" {
  depends_on = [resource.routeros_interface_wireless_security_profiles.test]
  security_profile=resource.routeros_interface_wireless_security_profiles.test.name
  mode="ap-bridge" 
  master_interface="wlan1" 
  name="wlan-guest"
  ssid="guests" 
  basic_rates_ag = ["6Mbps", "9Mbps"]
}
`, providerConfig)
}
