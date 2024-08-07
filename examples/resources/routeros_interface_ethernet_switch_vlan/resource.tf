resource "routeros_interface_ethernet_switch_vlan" "test" {
  switch               = "switch1"
  ports                = ["ether1"]
  vlan_id              = 10
  independent_learning = true
}