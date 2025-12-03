#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/list get [print show-ids]]
terraform import routeros_interface_list.list "*2000010"#Or you can import a resource using one of its attributes
terraform import routeros_interface_list.list "name=xxx"
