resource "routeros_routing_table" "test_table" {
  name = "to_ISP1"
  fib  = false
}