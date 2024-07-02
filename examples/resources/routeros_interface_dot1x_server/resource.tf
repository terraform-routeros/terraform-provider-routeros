resource "routeros_interface_dot1x_server" "ether2" {
  auth_types = ["mac-auth"]
  interface = "ether2"
}
