#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-server/option/sets/get [print show-ids]]
terraform import routeros_ip_dhcp_server_option_set.lan_option_set "*1"