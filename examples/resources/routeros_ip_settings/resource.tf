resource "routeros_ip_settings" "settings" {
  ipv4_multipath_hash_policy = "l3-inner"
}
