#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/traffic/flow/target get [print show-ids]]
terraform import routeros_ip_traffic_flow_target.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ip_traffic_flow_target.test "dst_address=xxx"
