resource "routeros_routing_bfd_configuration" "test" {
  interfaces = ["lo", "ether2"]
  vrf        = "main"
  forbid_bfd = true
}
