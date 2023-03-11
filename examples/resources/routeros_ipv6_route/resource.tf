resource "routeros_ipv6_route" "a_route" {
  dst_address = "::/0"
  gateway     = "2001:DB8:1000::1"
}