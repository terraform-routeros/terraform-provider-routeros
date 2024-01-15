resource "routeros_interface_ethernet_switch_rule" "test" {
  switch      = "switch1"
  ports       = ["ether1"]
  copy_to_cpu = true
}
