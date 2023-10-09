resource "routeros_user_manager_attribute" "mikrotik_wireless_comment" {
  name         = "Mikrotik-Wireless-Comment"
  packet_types = ["access-accept"]
  type_id      = 21
  value_type   = "string"
}
