# The ID can be found via API or the terminal
# :put [/interface/ethernet/switch/qos/map/ip get [print show-ids]]
terraform import routeros_interface_ethernet_switch_qos_map_ip.ef "*0"
# or by field:
terraform import routeros_interface_ethernet_switch_qos_map_ip.ef "dscp=46"
