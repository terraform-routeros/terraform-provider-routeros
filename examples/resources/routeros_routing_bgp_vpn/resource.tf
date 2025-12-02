resource "routeros_routing_bgp_vpn" "test" {
  disabled = false
  export {
    redistribute  = "connected"
    route_targets = ["1:1"]
  }
  import {
    route_targets = ["1:2"]
  }
  label_allocation_policy = "per-vrf"
  name                    = "bgp-mpls-vpn-test"
  route_distinguisher     = "1.2.3.4:1"
  vrf                     = "vrfTest1"
}
