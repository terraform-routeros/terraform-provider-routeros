resource "routeros_routing_rule" "test" {
  dst_address = "192.168.1.0/24"
  action      = "lookup-only-in-table"
  interface   = "ether1"
}