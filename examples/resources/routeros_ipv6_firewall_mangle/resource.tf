resource "routeros_ipv6_firewall_mangle" "rule" {
  action        = "change-mss"
  chain         = "forward"
  out_interface = "pppoe-out"
  protocol      = "tcp"
  tcp_flags     = "syn"
  new_mss       = "1130"
  tcp_mss       = "1301-65535"
}