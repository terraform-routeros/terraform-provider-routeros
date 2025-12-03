#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-server/network get [print show-ids]]
terraform import routeros_ip_dhcp_server_network.dhcp_server_network "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_ip_dhcp_server_network.dhcp_server_network "name=xxx"
