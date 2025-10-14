resource "routeros_routing_bgp_template" "test" {
  name = "test-template"
  as   = 65521
  input {
    limit_process_routes_ipv4 = 5
    limit_process_routes_ipv6 = 5
  }
  output {
    affinity             = "alone"
    keep_sent_attributes = true
    default_originate    = "never"
  }
  // save_to = "bgp.dump"
}
