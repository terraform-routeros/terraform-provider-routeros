resource "routeros_routing_id" "test" {
  name              = "router-id-test"
  router_id         = "10.10.10.10"
  select_dynamic_id = "any"
}
