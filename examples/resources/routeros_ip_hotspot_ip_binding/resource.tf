resource "routeros_ip_hotspot_ip_binding" "test" {
  address     = "0.0.0.1"
  comment     = "comment"
  mac_address = "00:00:00:00:01:10"
  to_address  = "0.0.0.2"
}
