resource "routeros_dns" "dns-server" {
  allow_remote_requests = true
  servers = [
    "2606:4700:4700::1111,1.1.1.1",
    "2606:4700:4700::1001,1.0.0.1",
  ]
}