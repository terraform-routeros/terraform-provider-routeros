#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ get [print show-ids]]
terraform import routeros_interface_lte_apn.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_lte_apn.test "name=xxx"
