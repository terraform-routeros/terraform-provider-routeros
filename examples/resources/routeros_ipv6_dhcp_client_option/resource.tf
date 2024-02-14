resource "routeros_ipv6_dhcp_client_option" "option" {
  name = "my-dhcp-option"
  code = 60
}
