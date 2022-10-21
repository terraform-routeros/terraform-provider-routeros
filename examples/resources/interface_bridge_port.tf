resource "routeros_interface_bridge_port" "bridge_port" {
    bridge    = "bridge"
    interface = "ether5"
    pvid      = "50"
}