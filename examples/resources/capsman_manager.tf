resource "routeros_capsman_manager" "manager" {
  enabled     = true
  certificate = "my-ssl-cert"
}
