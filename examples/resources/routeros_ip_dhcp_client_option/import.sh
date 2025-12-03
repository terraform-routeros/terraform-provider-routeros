#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-client/option get [print show-ids]]
terraform import routeros_ip_dhcp_client_option.option "*0"#Or you can import a resource using one of its attributes
terraform import routeros_ip_dhcp_client_option.option "name=xxx"
