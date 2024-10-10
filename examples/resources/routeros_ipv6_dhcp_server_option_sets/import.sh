#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ipv6/dhcp/server/option/sets get [print show-ids]]
terraform import routeros_ipv6_dhcp_server_option_sets.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ipv6_dhcp_server_option_sets.test "name=test-set"