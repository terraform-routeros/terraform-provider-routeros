resource "routeros_ipv6_firewall_nat" "rule" {
  action        = "masquerade"
  chain         = "srcnat"
  out_interface = "ether16"
}