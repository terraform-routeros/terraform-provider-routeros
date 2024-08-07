
resource "routeros_interface_lte_apn" "test" {
  name           = "apn1"
  apn            = "internet"
  authentication = "pap"
}