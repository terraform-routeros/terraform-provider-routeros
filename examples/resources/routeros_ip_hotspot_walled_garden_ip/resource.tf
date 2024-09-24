resource "routeros_ip_hotspot_walled_garden_ip" "test" {
  action           = "reject"
  dst_address      = "!0.0.0.0"
  dst_address_list = "dlist"
  dst_port         = "0-65535"
  protocol         = "tcp"
  src_address      = "0.0.0.0"
  src_address_list = "slist"
}
