#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/bridge/filter get [print show-ids]]
terraform import routeros_interface_bridge_filter.rule "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_interface_bridge_filter.rule "dst_address=224.0.0.251/32"