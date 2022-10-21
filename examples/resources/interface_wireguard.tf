resource "routeros_interface_wireguard" "test_wg_interface" {
    name        = "test_wg_interface"
    listen_port = "13231"
}