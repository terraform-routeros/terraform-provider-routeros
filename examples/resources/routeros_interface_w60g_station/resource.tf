resource "routeros_interface_w60g_station" "test" {
  name           = "wlan60-station-1"
  parent         = "wlan60-1"
  remote-address = "AA:AA:AA:AA:AA:AA"
  put-in-bridge  = "parent"
}
