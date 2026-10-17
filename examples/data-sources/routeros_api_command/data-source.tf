# Count the entries of an IPv6 firewall address list.
data "routeros_api_command" "blocklist_size" {
  path  = "/ipv6/firewall/address-list"
  words = ["count-only", "?list=blocklist"]
}

output "blocklist_size" {
  value = data.routeros_api_command.blocklist_size.result[0]["ret"]
}

# Read the rows of a menu that has no typed data source yet.
data "routeros_api_command" "ospf_neighbors" {
  path = "/routing/ospf/neighbor"
}
