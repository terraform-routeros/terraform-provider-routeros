#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/wifi/security/multi/passphrase get [print show-ids]]
terraform import routeros_wifi_security_multi_passphrase.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_wifi_security_multi_passphrase.test "comment=xxx"