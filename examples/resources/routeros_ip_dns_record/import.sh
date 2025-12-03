#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dns/static get [print show-ids]]
terraform import routeros_ip_dns_record.name_record "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_ip_dns_record.name_record "name=xxx"
