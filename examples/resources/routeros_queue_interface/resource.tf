# The per-interface row is auto-created by RouterOS; this adopts and configures it.
resource "routeros_queue_interface" "ether1" {
  interface = "ether1"
  queue     = "ethernet-default"
}
