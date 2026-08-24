# The per-port row is auto-created by the switch chip; this adopts and configures it.
resource "routeros_interface_ethernet_switch_qos_port" "sfp1" {
  name     = "sfp-sfpplus1"
  profile  = "5-Standard (BE/Default)"
  trust_l2 = "trust"
  trust_l3 = "trust"
}
