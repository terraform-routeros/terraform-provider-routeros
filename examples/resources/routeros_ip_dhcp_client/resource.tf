resource "routeros_ip_dhcp_client" "client" {
  interface = "bridge"
}