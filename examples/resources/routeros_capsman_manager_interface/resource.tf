resource "routeros_capsman_manager_interface" "test_manager_interface" {
  interface = "ether1"
  forbid    = true
}