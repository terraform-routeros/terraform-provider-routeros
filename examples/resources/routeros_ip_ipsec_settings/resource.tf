resource "routeros_ip_ipsec_settings" "test" {
  xauth_use_radius = true
  interim_update   = "60s"
}