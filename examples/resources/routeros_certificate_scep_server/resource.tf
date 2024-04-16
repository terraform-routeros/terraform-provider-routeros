resource "routeros_system_certificate" "example_root_ca" {
  name        = "example_root_ca"
  common_name = "Example Root CA"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  sign {
  }
}

resource "routeros_certificate_scep_server" "example_scep_server" {
  ca_cert    = routeros_system_certificate.example_root_ca.name
  path       = "/scep/example_scep_server"
  days_valid = 30
}
