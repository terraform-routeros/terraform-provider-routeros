resource "routeros_wifi_cap" "settings" {
  enabled              = true
  discovery_interfaces = ["bridge1"]
}
