resource "routeros_ip_dhcp_server_config" "settings" {
  accounting        = true
  interim_update    = "1m"
  radius_password   = "same-as-user"
  store_leases_disk = "10m"
}
