resource "routeros_ip_firewall_filter" "rule" {
  action      = "accept"
  chain       = "forward"
  src_address = "10.0.0.1"
  dst_address = "10.0.1.1"
  dst_port    = "443"
  protocol    = "tcp"
}