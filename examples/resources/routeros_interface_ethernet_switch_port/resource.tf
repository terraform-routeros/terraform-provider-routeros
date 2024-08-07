resource "routeros_interface_ethernet_switch_port" "test" {
  name      = "ether1"
  vlan_mode = "check"
}
