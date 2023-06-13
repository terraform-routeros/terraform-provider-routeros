resource "routeros_interface_bonding" "test_bonding" {
	name   = "test_bonding"
	slaves = "ether1"
}