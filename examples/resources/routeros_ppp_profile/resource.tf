resource "routeros_ppp_profile" "test" {
  name           = "ovpn"
  local_address  = "192.168.77.1"
  remote_address = "ovpn-pool"
  use_upnp       = "no"
}