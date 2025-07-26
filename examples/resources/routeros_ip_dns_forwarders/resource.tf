resource "routeros_ip_dns_forwarders" "test" {
  disabled        = true
  dns_servers     = ["1.1.1.1"]
  doh_servers     = ["2.2.2.2"]
  name            = "test"
  verify_doh_cert = "false"
}
