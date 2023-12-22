resource "routeros_wifi_configuration" "configuration1" {
  country = "Netherlands"
  manager = "capsman"
  mode    = "ap"
  name    = "configuration1"
  ssid    = "my-network"
}

resource "routeros_wifi_provisioning" "provisioning1" {
  action               = "create-enabled"
  master_configuration = routeros_wifi_configuration.configuration1.name
  name_format          = "cap1:"
  radio_mac            = "00:11:22:33:44:55"
}
