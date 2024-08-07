resource "routeros_routing_ospf_instance" "test_routing_ospf_instance" {
  name = "test_routing_ospf_instance"
}

resource "routeros_routing_ospf_area" "test_routing_ospf_area" {
  name     = "test_routing_ospf_area"
  instance = routeros_routing_ospf_instance.test_routing_ospf_instance.name
}

resource "routeros_routing_ospf_interface_template" "test_routing_ospf_interface_template" {
  area = routeros_routing_ospf_area.test_routing_ospf_area.name
}