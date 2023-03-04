resource "routeros_ip_pool" "pool" {
  name   = "my_ip_pool"
  ranges = "10.0.0.100-10.0.0.200"
}