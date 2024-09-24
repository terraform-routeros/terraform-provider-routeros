resource "routeros_ip_hotspot_profile" "test" {
  name       = "hsprof-1"
  login_by   = ["mac", "https", "trial"]
  use_radius = true
}
