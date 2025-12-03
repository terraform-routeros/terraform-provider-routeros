#The ID can be found via API or the terminal
#The command for the terminal is -> /routing/ospf/area/print show-ids
terraform import routeros_routing_ospf_area.test_routing_ospf_area "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_routing_ospf_area.test_routing_ospf_area "name=xxx"
