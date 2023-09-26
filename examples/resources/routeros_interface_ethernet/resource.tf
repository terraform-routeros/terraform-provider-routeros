resource "routeros_interface_ethernet" "test" {
  factory_name = "sfp-sfpplus8"
  name         = "swtich-eth0"
  mtu          = 9000
}
