#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ipv6/firewall/filter get [print show-ids]]
terraform import routeros_ipv6_firewall_filter.rule "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_ipv6_firewall_filter.rule "name=xxx"
