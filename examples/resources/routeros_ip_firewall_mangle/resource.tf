resource "routeros_ip_firewall_mangle" "rule" {
  action        = "change-mss"
  chain         = "forward"
  out-interface = "pppoe-out"
  protocol      = "tcp"
  tcp-flags     = "syn"
  new-mss       = "1130"
  tcp-mss       = "1301-65535"
}