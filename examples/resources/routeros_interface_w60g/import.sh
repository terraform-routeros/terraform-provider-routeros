#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/w60g get [print show-ids]]
terraform import routeros_interface_w60g.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_w60g.test "name=xxx"