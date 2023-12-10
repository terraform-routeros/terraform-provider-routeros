resource "routeros_wifi_aaa" "aaa1" {
  called_format   = "S"
  name            = "aaa1"
  password_format = ""
  username_format = "AA:AA:AA:AA:AA:AA"
}
