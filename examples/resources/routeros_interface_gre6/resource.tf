resource "routeros_interface_gre6" "gre_hq" {
  name           = "gre-hq-ipv6"
  remote_address = "2a02::2"
  disabled       = true
}
