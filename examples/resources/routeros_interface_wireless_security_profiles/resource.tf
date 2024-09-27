resource "routeros_interface_wireless_security_profiles" "test" {
  name                 = "test-profile"
  mode                 = "dynamic-keys"
  authentication_types = ["wpa-psk", "wpa2-psk"]
  wpa_pre_shared_key   = "wpa_psk_key"
  wpa2_pre_shared_key  = "wpa2_psk_key"
}