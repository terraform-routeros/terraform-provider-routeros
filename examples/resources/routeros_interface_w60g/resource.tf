resource "routeros_interface_w60g" "test" {
  make     = "wlan60-1"
  password = "put_your_safe_password_here"
  ssid     = "put_your_new_ssid_here"
  disabled = false
  mode     = "ap-bridge"
}
