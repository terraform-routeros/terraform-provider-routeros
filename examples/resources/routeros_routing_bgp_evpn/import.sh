#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/routing/bgp/evpn get [print show-ids]]
terraform import routeros_routing_bgp_evpn.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_routing_bgp_evpn.test "name=xxx"