locals {
  tls_service     = { "api-ssl" = 8729, "www-ssl" = 443 }
  disable_service = { "api" = 8728, "ftp" = 21, "telnet" = 23, "www" = 80 }
  enable_service  = { "ssh" = 22, "winbox" = 8291 }
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

# terraform state rm 'routeros_ip_service.tls["www-ssl"]'
# terraform import 'routeros_ip_service.tls["www-ssl"]' www-ssl
resource "routeros_ip_service" "tls" {
  for_each    = local.tls_service
  numbers     = each.key
  port        = each.value
  certificate = routeros_system_certificate.tls_cert.name
  tls_version = "only-1.2"
  disabled    = false
}

resource "routeros_ip_service" "disabled" {
  for_each = local.disable_service
  numbers  = each.key
  port     = each.value
  disabled = true
}

resource "routeros_ip_service" "enabled" {
  for_each = local.enable_service
  numbers  = each.key
  port     = each.value
  disabled = false
}