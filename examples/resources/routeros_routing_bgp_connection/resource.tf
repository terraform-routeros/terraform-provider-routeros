resource "routeros_routing_bgp_connection" "test" {
  name         = "neighbor-test"
  as           = "65550/5"
  as_override  = true
  add_path_out = "none"
  remote {
    address = "172.17.0.1"
    as      = "12345/5"
  }
  local {
    role = "ebgp"
  }
}