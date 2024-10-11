resource "routeros_ipv6_pool" "pool-0" {
  name          = "test-pool-0"
  prefix        = "2001:db8:40::/48"
  prefix_length = 64
}

resource "routeros_ipv6_dhcp_server" "test" {
  address_pool = routeros_ipv6_pool.pool-0.name
  interface    = "bridge"
  lease_time   = "1m"
  name         = "test-dhcpv6"
  preference   = 128
}
