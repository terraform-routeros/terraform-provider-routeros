resource "routeros_user_manager_limitation" "test" {
  name           = "test"
  download_limit = 1024
  upload_limit   = 1024
  uptime_limit   = "10d"
}

resource "routeros_user_manager_profile" "test" {
  name           = "test"
  name_for_users = "Test"
  price          = 0.02
}

resource "routeros_user_manager_profile_limitation" "weekend_night" {
  limitation = routeros_user_manager_limitation.test.name
  profile    = routeros_user_manager_profile.test.name
  from_time  = "0s"
  till_time  = "6h"
  weekdays = [
    "sunday",
    "saturday",
  ]
}
