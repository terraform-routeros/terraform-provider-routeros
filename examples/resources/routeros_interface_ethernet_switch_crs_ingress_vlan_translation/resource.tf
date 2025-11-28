resource "routeros_interface_ethernet_switch_crs_ingress_vlan_translation" "test" {
  ports                = ["ether1"]
  service_vlan_format  = "any"
  customer_vlan_format = "any"
  customer_vid         = 0
  new_customer_vid     = 100
  sa_learning          = true
}
