resource "routeros_ip_traffic_flow_ipfix" "test" {
  nat_events = true
  dst_port   = false
}
