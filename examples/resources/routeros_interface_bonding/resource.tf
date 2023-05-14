resource "routeros_interface_bonding" "test" {
  name   = "bonding-test"
  slaves = ["ether3", "ether4"]
}