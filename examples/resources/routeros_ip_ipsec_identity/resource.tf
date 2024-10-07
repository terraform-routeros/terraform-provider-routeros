resource "routeros_ip_ipsec_mode_config" "test" {
  name      = "NordVPN"
  responder = false
}

resource "routeros_ip_ipsec_peer" "test" {
  address       = "lv20.nordvpn.com"
  exchange_mode = "ike2"
  name          = "NordVPN"
}

resource "routeros_ip_ipsec_identity" "test" {
  auth-method     = "eap"
  certificate     = ""
  eap-methods     = "eap-mschapv2"
  generate-policy = "port-strict"
  mode-config     = routeros_ip_ipsec_mode_config.test.name
  peer            = routeros_ip_ipsec_peer.test.name
  username        = "support@mikrotik.com"
  password        = "secret"
}
