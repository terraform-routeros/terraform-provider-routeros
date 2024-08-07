resource "routeros_capsman_manager" "test_manager" {
  enabled        = true
  upgrade_policy = "require-same-version"
}