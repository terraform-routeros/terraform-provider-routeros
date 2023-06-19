resource "routeros_ip_cloud" "test" {
  ddns_enabled         = true
  update_time          = false
  ddns_update_interval = "11m"
}