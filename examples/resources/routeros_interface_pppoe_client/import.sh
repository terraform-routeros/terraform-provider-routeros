#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/pppoe-client get [print show-ids]]
terraform import routeros_interface_pppoe_client.test "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_interface_pppoe_client.test "name=xxx"
