#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ethernet/switch/port get [print show-ids]]
terraform import routeros_interface_ethernet_switch_port.test *1
#Or you can import a resource using one of its attributes
terraform import routeros_interface_ethernet_switch_port.test "name=xxx"
