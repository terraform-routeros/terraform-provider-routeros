resource "routeros_system_ntp_client" "test" {
  enabled = true
  mode    = "unicast"
  servers = ["146.59.35.38", "167.235.201.139"]
}