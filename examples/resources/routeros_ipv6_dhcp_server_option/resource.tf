resource "routeros_ipv6_dhcp_server_option" "test" {
  name  = "domain-search"
  code  = 24
  value = "0x07'example'0x05'local'0x00"
}