#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/routing/igmp/proxy/interface get [print show-ids]]
terraform import routeros_routing_igmp_proxy_interface.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_routing_igmp_proxy_interface.test "interface=xxx"