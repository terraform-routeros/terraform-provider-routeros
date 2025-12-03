#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/macvlan get [print show-ids]]
terraform import routeros_interface_macvlan.test "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_interface_macvlan.test "name=xxx"
