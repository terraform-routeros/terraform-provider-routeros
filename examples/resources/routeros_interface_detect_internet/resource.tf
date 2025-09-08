resource "routeros_interface_detect_internet" "test" {
  internet_interface_list = "all"
  wan_interface_list      = "all"
  lan_interface_list      = "all"
  detect_interface_list   = "all"
}
