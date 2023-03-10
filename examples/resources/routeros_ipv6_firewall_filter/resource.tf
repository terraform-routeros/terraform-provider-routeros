resource "routeros_ipv6_firewall_filter" "rule" {
  action      = "accept"
  chain       = "forward"
  src_address = "2001:DB8:1000::1"
  dst_address = "2001:DB8:2000::1"
  dst_port    = "443"
  protocol    = "tcp"
}