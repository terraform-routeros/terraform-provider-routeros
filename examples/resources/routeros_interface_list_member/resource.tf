resource "routeros_interface_list_member" "list_member" {
  interface = "ether1"
  list      = "my-list"
}