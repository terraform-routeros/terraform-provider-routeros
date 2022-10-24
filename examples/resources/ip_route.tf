resource "routeros_ip_route" "a_route" {
  dst_address = "0.0.0.0/0"
  gateway     = "10.0.0.1"
}