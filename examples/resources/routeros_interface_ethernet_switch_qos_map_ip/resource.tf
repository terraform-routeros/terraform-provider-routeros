resource "routeros_interface_ethernet_switch_qos_map_ip" "ef" {
  dscp    = 46
  profile = "1-Real-Time (EF/VoIP/Video)"
}
