#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/gre get [print show-ids]]
terraform import routeros_interface_gre.gre_hq "*1"
#Or you can import a resource using one of its attributes
terraform import routeros_interface_gre.gre_hq "name=xxx"
