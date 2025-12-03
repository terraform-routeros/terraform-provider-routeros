#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/bridge/port get [print show-ids]]
terraform import routeros_interface_bridge_port.bridge_port "*0"#Or you can import a resource using one of its attributes
terraform import routeros_interface_bridge_port.bridge_port "name=xxx"
