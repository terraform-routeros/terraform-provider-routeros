resource "routeros_interface_vxlan" "test" {
  name = "vxlan1-test"
  vni  = 10
}
