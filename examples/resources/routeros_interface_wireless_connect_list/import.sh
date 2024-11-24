#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wireless/connect-list get [print show-ids]]
terraform import routeros_interface_wireless_connect_list.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_wireless_connect_list.test "name=xxx"