resource "routeros_ipv6_pool" "test" {
  name          = "test-pool"
  prefix        = "2001:db8:12::/48"
  prefix_length = 64
}
