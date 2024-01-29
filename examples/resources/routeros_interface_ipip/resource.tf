resource "routeros_interface_ipip" "ipip_hq" {
  name           = "ipip-hq-1"
  remote_address = "10.77.3.26"
  disabled       = true
}
