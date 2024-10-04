resource "routeros_ip_ipsec_mode_config" "test" {
  name          = "test-cfg"
  address       = "1.2.3.4"
  split_include = ["0.0.0.0/0"]
  split_dns     = ["1.1.1.1"]
}
