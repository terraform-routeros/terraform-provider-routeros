resource "routeros_interface_ethernet_switch_crs_egress_vlan_tag" "test" {
  vlan_id      = 100
  tagged_ports = ["ether1", "ether2"]
}
