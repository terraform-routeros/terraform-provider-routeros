# The ID can be found via API or the terminal
# :put [/interface/ethernet/switch/qos/profile get [print show-ids]]
terraform import routeros_interface_ethernet_switch_qos_profile.realtime "*1"
# or by name:
terraform import routeros_interface_ethernet_switch_qos_profile.realtime "name=1-Real-Time (EF/VoIP/Video)"
