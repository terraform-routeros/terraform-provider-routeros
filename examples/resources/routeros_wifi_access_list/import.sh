#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wifi/access-list get [print show-ids]]
terraform import routeros_wifi_access_list.radius '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_wifi_access_list.radius "name=xxx"
