resource "routeros_ip_dhcp_server_lease" "dhcp_lease" {
  address     = "10.0.0.2"
  mac_address = "AA:BB:CC:DD:11:22"
}