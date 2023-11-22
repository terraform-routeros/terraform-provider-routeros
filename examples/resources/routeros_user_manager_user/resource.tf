resource "routeros_user_manager_attribute" "mikrotik_wireless_comment" {
  name       = "Mikrotik-Wireless-Comment"
  type_id    = 21
  value_type = "string"
}

resource "routeros_user_manager_attribute" "mikrotik_wireless_vlanid" {
  name       = "Mikrotik-Wireless-VLANID"
  type_id    = 26
  value_type = "uint32"
}

resource "routeros_user_manager_user_group" "test" {
  name = "test"
}

resource "routeros_user_manager_user" "test" {
  attributes = [
    "${routeros_user_manager_attribute.mikrotik_wireless_comment.name}:Test Group",
    "${routeros_user_manager_attribute.mikrotik_wireless_vlanid.name}:100",
  ]
  group    = routeros_user_manager_user_group.test.name
  name     = "test"
  password = "password"
}
