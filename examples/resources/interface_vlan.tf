resource "routeros_interface_vlan" "interface_vlan" {
    interface = "bridge"
    name      = "VLAN_TEST"
    vlan_id   = 50
}