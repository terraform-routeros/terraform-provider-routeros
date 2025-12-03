#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-client get [print show-ids]]
terraform import routeros_ip_dhcp_client.client "*0"#Or you can import a resource using one of its attributes
terraform import routeros_ip_dhcp_client.client "name=xxx"
