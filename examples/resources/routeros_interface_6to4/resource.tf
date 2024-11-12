resource "routeros_interface_6to4" "test" {
  name      = "6to4-tunnel1"
  keepalive = "10,10"
}
