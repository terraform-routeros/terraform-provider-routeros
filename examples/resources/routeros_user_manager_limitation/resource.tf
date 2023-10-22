resource "routeros_user_manager_limitation" "test" {
  name           = "test"
  download_limit = 1024
  upload_limit   = 1024
  uptime_limit   = "10d"
}
