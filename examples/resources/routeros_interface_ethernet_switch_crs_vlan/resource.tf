resource "routeros_interface_ethernet_switch_crs_vlan" "test" {
  ports   = ["ether1"]
  vlan_id = 10
}
