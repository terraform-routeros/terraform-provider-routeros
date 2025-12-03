#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/vrrp get [print show-ids]]
terraform import routeros_interface_vrrp.interface_vrrp "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_interface_vrrp.interface_vrrp "name=xxx"
