#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/ipsec/mode/config get [print show-ids]]
terraform import routeros_ip_ipsec_mode_config.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ip_ipsec_mode_config.test "address=1.2.3.4"