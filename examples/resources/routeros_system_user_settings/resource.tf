resource "routeros_system_user_settings" "settings" {
  minimum_categories      = 2
  minimum_password_length = 8
}
