resource "routeros_ip_traffic_flow_target" "test" {
  dst_address = "192.168.0.2"
  port        = 2055
  version     = "9"
}
