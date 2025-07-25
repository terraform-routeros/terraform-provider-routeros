resource "routeros_routing_igmp_proxy_interface" "test" {
  alternative_subnets = ["0.0.0.1/32", "0.0.0.2/32"]
  disabled            = true
  interface           = "lo"
  threshold           = 5
}
