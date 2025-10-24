resource "routeros_routing_ospf_area_range" "test" {
  area      = routeros_routing_ospf_area.ospf-area-1.name
  advertise = true
  prefix    = "::/64"
  disabled  = true
}
