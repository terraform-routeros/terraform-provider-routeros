resource "routeros_interface_ethernet_switch_qos_profile" "realtime" {
  name          = "1-Real-Time (EF/VoIP/Video)"
  dscp          = 46
  traffic_class = 7
}
