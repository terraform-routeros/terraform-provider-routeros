resource "routeros_ip_upnp_interfaces" "test" {
  interface          = "ether1"
  type               = "external"
  forced_external_ip = "0.0.0.0"
}