resource "routeros_interface_ethernet_switch_port_isolation" "test" {
  name                = "ether1"
  forwarding_override = "ether1"
}
