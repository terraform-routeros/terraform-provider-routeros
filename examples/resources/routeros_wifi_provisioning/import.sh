#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wifi/provisioning get [print show-ids]]
terraform import routeros_wifi_provisioning.provisioning1 '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_wifi_provisioning.provisioning1 "name=xxx"
