resource "routeros_interface_ovpn_server" "user1" {
  name       = "ovpn-in1"
  user       = "user1"
  depends_on = [routeros_ovpn_server.server]
}
