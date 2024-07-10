resource "routeros_interface_dot1x_client" "ether2" {
  eap_methods = ["eap-peap", "eap-mschapv2"]
  identity = "router"
  interface = "ether2"
}
