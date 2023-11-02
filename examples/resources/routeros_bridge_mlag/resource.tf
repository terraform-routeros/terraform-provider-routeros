resource "routeros_bridge_mlag" "mlag" {
  bridge    = "bridge1"
  peer_port = "stack-link"
}