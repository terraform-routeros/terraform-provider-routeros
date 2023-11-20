resource "routeros_user_manager_router" "test" {
  address       = "127.0.0.1"
  name          = "test"
  shared_secret = "password"
}
