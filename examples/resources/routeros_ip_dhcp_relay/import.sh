#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-relay get [print show-ids]]
terraform import routeros_ip_dhcp_relay.relay "*0"#Or you can import a resource using one of its attributes
terraform import routeros_ip_dhcp_relay.relay "name=xxx"
