resource "routeros_interface_bridge" "bridge" {
  name           = "bridge"
  vlan_filtering = true
}

# RouterOS 7.22+ MLAG example.
resource "routeros_interface_bridge" "mlag_bridge" {
  name           = "bridge1"
  vlan_filtering = true
  mlag_peer_port = "stack-link"
  mlag_heartbeat = "5s"
  mlag_priority  = 100
}
