resource "routeros_ipv6_dhcp_client" "inet_provider" {
  pool_name         = "pub-add-pool"
  interface         = "ether1"
  add-default-route = true
  pool_prefix_length = 64
  request   = ["prefix"]
  disabled  = false
}

resource "routeros_ipv6_dhcp_client" "client" {
  pool_name         = "pub-add-pool"
  interface         = "ether1"
  add-default-route = true
  pool_prefix_length = 64
  request            = ["prefix"]
}
