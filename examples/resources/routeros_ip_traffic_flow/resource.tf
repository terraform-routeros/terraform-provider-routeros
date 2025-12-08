resource "routeros_ip_traffic_flow" "test" {
  packet_sampling   = true
  sampling_interval = 2222
  sampling_space    = 1111
}
