resource "routeros_ip_firewall_layer7_protocol" "test" {
  name   = "rdp"
  regexp = "rdpdr.*cliprdr.*rdpsnd"
}
