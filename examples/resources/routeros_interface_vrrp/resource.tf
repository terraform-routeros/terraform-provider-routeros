resource "routeros_interface_vrrp" "interface_vrrp" {
  interface = "bridge"
  name      = "lan_vrrp"
}