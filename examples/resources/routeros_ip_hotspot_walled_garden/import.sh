#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/hotspot/walled-garden get [print show-ids]]
terraform import routeros_ip_hotspot_walled_garden.test *3#Or you can import a resource using one of its attributes
terraform import routeros_ip_hotspot_walled_garden.test "name=xxx"
