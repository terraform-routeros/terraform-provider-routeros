resource "routeros_ip_dhcp_client_option" "option" {
  name = "my-dhcp-option"
  code = 60
}