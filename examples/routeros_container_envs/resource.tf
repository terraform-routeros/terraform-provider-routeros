resource "routeros_container_envs" "test_envs" {
  name  = "test_envs"
  key   = "TZ"
  value = "UTC"
}
