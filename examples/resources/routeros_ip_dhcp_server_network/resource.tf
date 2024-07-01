resource "routeros_ip_dhcp_server_network" "dhcp_server_network" {
  address    = "10.0.0.0/24"
  gateway    = "10.0.0.1"
  dns_server = ["1.1.1.1"]
}