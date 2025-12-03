#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/radius get [print show-ids]]
terraform import routeros_radius.user_manager *1
#Or you can import a resource using one of its attributes
terraform import routeros_radius.user_manager "name=xxx"
