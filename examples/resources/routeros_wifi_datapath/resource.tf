resource "routeros_wifi_datapath" "datapath1" {
  name             = "datapath1"
  bridge           = "bridge1"
  client_isolation = false
}
