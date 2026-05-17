# RouterOS versions before 7.22.
resource "routeros_bridge_mlag" "mlag" {
  bridge    = "bridge1"
  peer_port = "stack-link"
}
