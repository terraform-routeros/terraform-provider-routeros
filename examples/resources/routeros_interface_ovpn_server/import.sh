#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ovpn-server get [print show-ids]]
terraform import routeros_interface_ovpn_server.user1 *29#Or you can import a resource using one of its attributes
terraform import routeros_interface_ovpn_server.user1 "name=xxx"
