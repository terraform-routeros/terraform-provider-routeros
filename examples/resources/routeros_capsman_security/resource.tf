resource "routeros_capsman_security" "test_security" {
  name                  = "test_security"
  comment               = "test_security"
  authentication_types  = ["wpa-psk", "wpa-eap", "wpa2-psk"]
  disable_pmkid         = true
  eap_methods           = "eap-tls,passthrough"
  eap_radius_accounting = true
  encryption            = ["tkip", "aes-ccm"]
  group_encryption      = "aes-ccm"
  group_key_update      = "1h"
  passphrase            = "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDE"
  tls_certificate       = "none"
  tls_mode              = "verify-certificate"
}