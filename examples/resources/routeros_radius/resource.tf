resource "routeros_radius" "user_manager" {
  address = "127.0.0.1"
  service = ["ppp", "login"]
}
