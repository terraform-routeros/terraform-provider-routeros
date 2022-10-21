resource "routeros_ip_dhcp_server" "server" {
  address_pool = "my_address_pool"
  interface    = "bridge"
  name         = "bridge_dhcp"
}