#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ethernet/switch/ingress-vlan-translation get [print show-ids]]
terraform import routeros_interface_ethernet_switch_crs_ingress_vlan_translation.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_ethernet_switch_crs_ingress_vlan_translation.test "name=xxx"