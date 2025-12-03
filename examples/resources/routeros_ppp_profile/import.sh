#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ppp/profile get [print show-ids]]
terraform import routeros_ppp_profile.test *6#Or you can import a resource using one of its attributes
terraform import routeros_ppp_profile.test "name=xxx"
