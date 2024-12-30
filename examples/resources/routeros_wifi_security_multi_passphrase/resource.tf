resource "routeros_wifi_security_multi_passphrase" "test" {
  group      = "gr-123"
  passphrase = data.vault_kv_secret_v2.wifi_security.data["test"]
}
