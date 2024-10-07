#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/ipsec/policy/group get [print show-ids]]
terraform import routeros_ip_ipsec_policy_group.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ip_ipsec_policy_group.test "name=test-group"