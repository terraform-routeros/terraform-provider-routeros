resource "routeros_interface_gre" "gre_hq" {
  name           = "gre-hq-1"
  remote_address = "10.77.3.26"
  disabled       = true
}
