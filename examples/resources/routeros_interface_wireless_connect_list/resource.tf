resource "routeros_interface_wireless_connect_list" "test" {
  interface        = "wlan0"
  security_profile = "test-secp"
}
