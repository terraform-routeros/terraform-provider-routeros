# The eight tx-manager queues (traffic-class 0..7) are auto-created; this adopts one.
resource "routeros_interface_ethernet_switch_qos_tx_manager_queue" "tc6" {
  traffic_class = 6
  schedule      = "high-priority-group"
  weight        = 16
}
