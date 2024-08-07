resource "routeros_interface_macvlan" "test" {
  interface = "ether1"
  name      = "macvlan1"
  disabled  = false
}