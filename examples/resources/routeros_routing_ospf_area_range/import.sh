#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/routing/ospf/area/range get [print show-ids]]
terraform import routeros_routing_ospf_area_range.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_routing_ospf_area_range.test "comment=xxx"