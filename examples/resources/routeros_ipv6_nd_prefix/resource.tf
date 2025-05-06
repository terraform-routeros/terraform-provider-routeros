resource "routeros_ipv6_nd_prefix" "test" {
  prefix             = "fd55::/64"
  interface          = "ether1"
  preferred_lifetime = "6d24h"
}