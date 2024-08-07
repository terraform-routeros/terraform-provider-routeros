
resource "routeros_interface_lte" "test" {
  allow_roaming = false
  apn_profiles  = "default"
  band          = []
  default_name  = "lte1"
  disabled      = false
  mtu           = "1500"
  name          = "lte1"
  network_mode  = ["3g", "lte"]
  sms_protocol  = null
}