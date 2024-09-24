resource "routeros_ip_hotspot_walled_garden" "test" {
  action   = "deny"
  dst_host = "1.2.3.4"
  dst_port = "!443"
}