locals {
  tls_service     = ["api-ssl", "www-ssl"]
  disable_service = ["api", "ftp", "telnet", "www"]
  enable_service  = ["ssh", "winbox"]
}

resource "routeros_system_certificate" "tls_cert" {
  name        = "tls-cert"
  common_name = "Mikrotik Router"
  days_valid  = 3650
  key_usage   = ["key-cert-sign", "crl-sign", "digital-signature", "key-agreement", "tls-server"]
  key_size    = "prime256v1"
  sign {
  }
}

resource "routeros_ip_service" "tls" {
  for_each    = toset(local.tls_service)
  numbers     = each.key
  certificate = routeros_system_certificate.tls_cert.name
  tls_version = "only-1.2"
  disabled    = false
}

resource "routeros_ip_service" "disabled" {
  for_each = toset(local.disable_service)
  numbers  = each.key
  disabled = true
}

resource "routeros_ip_service" "enabled" {
  for_each = toset(local.enable_service)
  numbers  = each.key
  disabled = false
}