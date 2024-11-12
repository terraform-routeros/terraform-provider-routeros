#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/6to4 get [print show-ids]]
terraform import routeros_interface_6to4.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_6to4.test "name=6to4-tunnel1"