resource "routeros_wifi_access_list" "radius" {
  action            = "query-radius"
  comment           = "RADIUS"
  radius_accounting = true
}
