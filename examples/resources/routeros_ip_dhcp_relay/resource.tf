resource "routeros_ip_dhcp_relay" "relay" {
  name        = "test relay"
  interface   = "ether1"
  dhcp_server = "0.0.0.1"
}