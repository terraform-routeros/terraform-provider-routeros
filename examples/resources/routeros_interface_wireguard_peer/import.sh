#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wireguard/peers get [print show-ids]]
terraform import routeros_interface_wireguard_peer.wg_peer "*0"#Or you can import a resource using one of its attributes
terraform import routeros_interface_wireguard_peer.wg_peer "name=xxx"
