resource "routeros_interface_wireless_access_list" "test" {
  signal_range = "-100..100"
  time         = "3h3m-5h,mon,tue,wed,thu,fri"
  mac_address  = "00:AA:BB:CC:DD:EE"
}