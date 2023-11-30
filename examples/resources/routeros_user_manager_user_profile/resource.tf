resource "routeros_user_manager_profile" "test" {
  name = "test"
}

resource "routeros_user_manager_user" "test" {
  name = "test"
}

resource "routeros_user_manager_user_profile" "test" {
  profile = routeros_user_manager_profile.test.name
  user    = routeros_user_manager_user.test.name
}
