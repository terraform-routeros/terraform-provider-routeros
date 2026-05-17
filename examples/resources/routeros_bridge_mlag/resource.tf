# Legacy MLAG resource for RouterOS versions before 7.22.
# RouterOS 7.22+ manages MLAG with routeros_interface_bridge.mlag_* attributes.
resource "routeros_bridge_mlag" "mlag" {
  bridge    = "bridge1"
  peer_port = "stack-link"
}
