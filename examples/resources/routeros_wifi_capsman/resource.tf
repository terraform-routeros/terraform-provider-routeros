resource "routeros_wifi_capsman" "settings" {
  enabled        = true
  interfaces     = ["bridge1"]
  upgrade_policy = "suggest-same-version"
}
