resource "routeros_ppp_secret" "test" {
  name     = "user-test"
  password = "123"
  profile  = "default"
}
