resource "routeros_ipv6_address" "ipv6_address" {
  address   = "fd55::1/64"
  interface = "ether1"
  disabled  = false
}
