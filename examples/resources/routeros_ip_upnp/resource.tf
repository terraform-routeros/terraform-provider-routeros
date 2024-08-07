resource "routeros_ip_upnp" "test" {
  allow_disable_external_interface = true
  enabled                          = true
  show_dummy_rule                  = true
}