variable "wlan_2ghz_disabled" {
  type    = bool
  default = false
}

resource "routeros_interface_wireless" "wlan-2ghz" {
  name     = "wlan1"
  disabled = var.wlan_2ghz_disabled
}

resource "routeros_interface_wireless_security_profiles" "test" {
  name                 = "test-profile"
  mode                 = "dynamic-keys"
  authentication_types = ["wpa-psk", "wpa2-psk"]
  wpa_pre_shared_key   = "wpa_psk_key"
  wpa2_pre_shared_key  = "wpa2_psk_key"
}

resource "routeros_interface_wireless" "test" {
  depends_on       = [resource.routeros_interface_wireless_security_profiles.test]
  security_profile = resource.routeros_interface_wireless_security_profiles.test.name
  mode             = "ap-bridge"
  master_interface = resource.routeros_interface_wireless.wlan-2ghz.name
  name             = "wlan-guest"
  ssid             = "guests"
  basic_rates_ag   = ["6Mbps", "9Mbps"]
}
