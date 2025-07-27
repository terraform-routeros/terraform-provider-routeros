
resource "routeros_ip_dhcp_server_option_matcher" "dhcp1_ip_by_vendor_class" {
  name         = "dhcp1_ip_by_vendor_class"
  server       = "dhcp1"
  address_pool = "pool1"

  code          = 60 # Vendor Class Identifier
  value         = "android-dhcp-11"
  matching_type = "exact"
}