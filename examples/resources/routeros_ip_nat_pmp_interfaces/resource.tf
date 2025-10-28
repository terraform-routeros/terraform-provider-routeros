resource "routeros_ip_nat_pmp_interfaces" "test" {
  interface = "ether1"
  type      = "external"
  forced_ip = "0.0.0.0"
}