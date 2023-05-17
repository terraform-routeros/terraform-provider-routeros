resource "routeros_ipv6_firewall_addr_list" "example_list" {
  address = "123:dead:beaf::/64"
  list    = "Example List"
}