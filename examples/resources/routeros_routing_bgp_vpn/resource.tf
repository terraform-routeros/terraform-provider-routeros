resource "routeros_routing_bgp_instance" "test" {
  as   = "65000"
  name = "bgp-instance-1"
}

resource "routeros_routing_bgp_vpn" "test" {
  disabled = false
  export {
    redistribute  = "connected"
    route_targets = ["1:1"]
  }
  import {
    route_targets = ["1:2"]
  }
  instance                = routeros_routing_bgp_instance.test.name
  label_allocation_policy = "per-vrf"
  name                    = "bgp-mpls-vpn-test"
  route_distinguisher     = "1.2.3.4:1"
  vrf                     = "vrfTest1"
}
