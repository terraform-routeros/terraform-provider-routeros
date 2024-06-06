#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/vrf get [print show-ids]]
terraform import routeros_ip_vrf.test_vrf_a "*0"
# or
terraform import routeros_ip_vrf.test_vrf_a "vrf_1"
# or
terraform import routeros_ip_vrf.test_vrf_a `"comment=Custom routing"`