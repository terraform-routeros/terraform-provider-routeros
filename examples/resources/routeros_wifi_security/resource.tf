resource "routeros_wifi_security" "security1" {
  name                 = "security1"
  authentication_types = ["wpa2-psk", "wpa3-psk"]
  ft                   = true
  ft_preserve_vlanid   = true
  passphrase           = "password"
}
