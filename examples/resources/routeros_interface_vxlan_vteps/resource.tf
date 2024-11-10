resource "routeros_interface_vxlan" "test-2" {
  name = "vxlan2-test"
  vni  = 11
}

resource "routeros_interface_vxlan_vteps" "test" {
  interface = routeros_interface_vxlan.test-2.name
  remote_ip = "192.168.10.10"
}
