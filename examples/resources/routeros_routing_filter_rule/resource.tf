resource "routeros_routing_filter_rule" "test" {
  chain    = "testChain"
  rule     = "if (dst in 192.168.1.0/24 && dst-len>24) {set distance +1; accept} else {set distance -1; accept}"
  comment  = "comment"
  disabled = true
}