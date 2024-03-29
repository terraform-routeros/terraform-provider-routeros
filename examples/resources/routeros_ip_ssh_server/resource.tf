resource "routeros_ip_ssh_server" "test" {
  strong_crypto      = true
  forwarding_enabled = "local"
  host_key_size      = 4096
}
