#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wifi/aaa get [print show-ids]]
terraform import routeros_wifi_aaa.aaa1 '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_wifi_aaa.aaa1 "name=xxx"
