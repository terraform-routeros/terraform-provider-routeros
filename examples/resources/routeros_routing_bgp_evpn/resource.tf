resource "routeros_routing_bgp_instance" "test" {
  as   = "65000"
  name = "bgp-instance-1"
}

resource "routeros_routing_bgp_evpn" "test" {
  disabled = false
  export {
    route_targets = ["1010:1010"]
  }
  import {
    route_targets = ["1010:1010"]
  }
  instance = resource.routeros_routing_bgp_instance.test.name
  name     = "bgp-evpn-1"
  vni      = 1010
}
