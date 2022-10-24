resource "routeros_capsman_security" "security" {
  name       = "default_security"
  passphrase = "my-super-secret-passphrase"
  encryption = "wpa2"
}
