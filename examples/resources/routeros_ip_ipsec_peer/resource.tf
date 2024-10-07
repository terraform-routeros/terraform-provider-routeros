resource "routeros_ip_ipsec_peer" "test" {
  address       = "lv20.nordvpn.com"
  exchange_mode = "ike2"
  name          = "NordVPN"
}
