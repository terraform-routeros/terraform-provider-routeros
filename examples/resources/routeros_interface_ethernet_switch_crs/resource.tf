resource "routeros_interface_ethernet_switch_crs" "sw0" {
  switch_id                                               = 0
  name                                                    = "new switch"
  drop_if_invalid_or_src_port_not_member_of_vlan_on_ports = ["ether1", "ether2", "ether3"]
}
