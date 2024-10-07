resource "routeros_ip_ipsec_key" "test" {
  name     = "test-key"
  key_size = 2048
}
