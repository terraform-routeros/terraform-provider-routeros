resource "routeros_system_keymat_provider" "hkdf" {
  name        = "hkdf-keymat"
  type        = "hkdf"
  hkdf_hash   = "sha256"
  hkdf_secret = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
}