resource "routeros_interface_bridge" "bridge" {
  name           = "bridge"
  vlan_filtering = true
}