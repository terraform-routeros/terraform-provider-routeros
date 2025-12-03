#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/eoip get [print show-ids]]
terraform import routeros_interface_eoip.eoip_tunnel1 *B
#Or you can import a resource using one of its attributes
terraform import routeros_interface_eoip.eoip_tunnel1 "name=xxx"
