resource "routeros_queue_simple" "test" {
  name      = "server"
  target    = ["10.1.1.1/32"]
  max_limit = "0/0"
}
