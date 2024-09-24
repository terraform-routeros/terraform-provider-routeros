resource "routeros_ip_hotspot_user_profile" "test" {
  add_mac_cookie     = true
  address_list       = "list-1"
  idle_timeout       = "none"
  keepalive_timeout  = "2m"
  mac_cookie_timeout = "3d"
  name               = "new-profile"
  shared_users       = 3
  status_autorefresh = "2m"
  transparent_proxy  = true
  advertise          = true
}
