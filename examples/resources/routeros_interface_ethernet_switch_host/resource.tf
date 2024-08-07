resource "routeros_interface_ethernet_switch_host" "test" {
  switch      = "switch1"
  mac_address = "00:00:00:00:00:00"
  ports       = ["ether1"]
  mirror      = true
}