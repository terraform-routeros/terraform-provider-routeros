resource "routeros_ip_ipsec_profile" "test" {
  name           = "test-profile"
  hash_algorithm = "sha256"
  enc_algorithm  = ["aes-192", "aes-256"]
  dh_group       = ["ecp384", "ecp521"]
  nat_traversal  = false
}