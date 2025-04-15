resource "routeros_queue_tree" "test" {
  name        = "server"
  parent      = "global"
  max_limit   = "10M"
  packet_mark = ["pmark-server"]
}
