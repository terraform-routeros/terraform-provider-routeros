resource "routeros_interface_eoip" "eoip_tunnel1" {
  name           = "eoip-tunnel1"
  local_address  = "192.168.88.1"
  remote_address = "192.168.88.2"
  disabled       = true
}
