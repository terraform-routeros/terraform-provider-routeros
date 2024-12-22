resource "routeros_interface_pppoe_server" "test" {
  comment  = "comment"
  disabled = true
  name     = "pppoe-in1"
  user     = ""
  service  = ""
}
