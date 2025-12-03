#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/nat-pmp/interfaces get [print show-ids]]
terraform import routeros_ip_nat_pmp_interfaces.test '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_ip_nat_pmp_interfaces.test "name=xxx"
