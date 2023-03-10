resource "routeros_ip_firewall_addr_list" "example_list" {
  address = "1.1.1.1"
  list    = "Example List"
}