resource "routeros_capsman_aaa" "test_3a" {
  called_format = "ssid"
  mac_mode      = "as-username-and-password"
}