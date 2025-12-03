#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/routing/bgp/instance get [print show-ids]]
terraform import routeros_routing_bgp_instance.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_routing_bgp_instance.test "name=xxx"#Or you can import a resource using one of its attributes
terraform import routeros_routing_bgp_instance.test
routeros_routing_bgp_instance.test "name=xxx"
