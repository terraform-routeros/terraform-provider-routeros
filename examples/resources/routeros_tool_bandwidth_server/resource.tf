resource "routeros_tool_bandwidth_server" "test" {
  enabled                 = true
  authenticate            = false
  max_sessions            = 100
  allocate_udp_ports_from = 2000
}
