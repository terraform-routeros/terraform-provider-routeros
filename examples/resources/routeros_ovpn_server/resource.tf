resource "routeros_ip_pool" "ovpn-pool" {
  name   = "ovpn-pool"
  ranges = ["192.168.77.2-192.168.77.254"]
}

resource "routeros_system_certificate" "ovpn_ca" {
  name        = "OpenVPN-Root-CA"
  common_name = "OpenVPN Root CA"
  key_size    = "prime256v1"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  sign {
  }
}

resource "routeros_system_certificate" "ovpn_server_crt" {
  name        = "OpenVPN-Server-Certificate"
  common_name = "Mikrotik OpenVPN"
  key_size    = "prime256v1"
  key_usage   = ["digital-signature", "key-encipherment", "tls-server"]
  sign {
    ca = routeros_system_certificate.ovpn_ca.name
  }
}

resource "routeros_ppp_profile" "test" {
  name           = "ovpn"
  local_address  = "192.168.77.1"
  remote_address = "ovpn-pool"
  use_upnp       = "no"
}

resource "routeros_ppp_secret" "test" {
  name     = "user-test"
  password = "123"
  profile  = routeros_ppp_profile.test.name
}

resource "routeros_ovpn_server" "server" {
  enabled         = true
  certificate     = routeros_system_certificate.ovpn_server_crt.name
  auth            = ["sha256", "sha512"]
  tls_version     = "only-1.2"
  default_profile = routeros_ppp_profile.test.name
}

# The resource should be created only after the OpenVPN server is enabled!
resource "routeros_interface_ovpn_server" "user1" {
  name       = "ovpn-in1"
  user       = "user1"
  depends_on = [routeros_ovpn_server.server]
}
