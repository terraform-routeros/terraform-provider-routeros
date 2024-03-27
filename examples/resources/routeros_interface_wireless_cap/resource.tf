resource "routeros_interface_wireless_cap" "settings" {
  discovery_interfaces = ["bridge1"]
  enabled              = true
  interfaces           = ["wlan1", "wlan2"]
}
