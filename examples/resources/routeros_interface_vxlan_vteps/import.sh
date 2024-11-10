#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/vxlan/vteps get [print show-ids]]
terraform import routeros_interface_vxlan_vteps.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_vxlan_vteps.test "interface=vxlan2-test"