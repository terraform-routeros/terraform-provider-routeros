resource "routeros_interface_veth" "test" {
  name    = "veth-test"
  address = "192.168.120.2/24"
  gateway = "192.168.120.1"
  comment = "Virtual interface"
}
