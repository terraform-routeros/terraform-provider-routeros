resource "routeros_ipv6_dhcp_server_option" "domain-search" {
  name  = "domain-search"
  code  = 24
  value = "0x07'example'0x05'local'0x00"
}

resource "routeros_ipv6_dhcp_server_option_sets" "test" {
  name    = "test-set"
  options = [routeros_ipv6_dhcp_server_option.domain-search.name]
}
