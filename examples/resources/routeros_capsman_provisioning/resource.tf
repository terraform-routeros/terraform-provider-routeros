resource "routeros_capsman_configuration" "test_configuration" {
  name = "cfg1"
}

resource "routeros_capsman_provisioning" "test_provisioning" {
  master_configuration = "cfg1"
  action               = "create-disabled"
  name_prefix          = "cap-"

  depends_on = [
    routeros_capsman_configuration.test_configuration,
  ]
}