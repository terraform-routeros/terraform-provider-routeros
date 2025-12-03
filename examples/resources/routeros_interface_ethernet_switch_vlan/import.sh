#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ethernet/switch/vlan get [print show-ids]]
terraform import routeros_interface_ethernet_switch_vlan.test *0
#Or you can import a resource using one of its attributes
terraform import routeros_interface_ethernet_switch_vlan.test "name=xxx"
