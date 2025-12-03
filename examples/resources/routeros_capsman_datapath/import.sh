#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/caps-man/datapath get [print show-ids]]
terraform import routeros_capsman_datapath.test_datapath "*1"
#Or you can import a resource using one of its attributes
terraform import routeros_capsman_datapath.test_datapath "name=xxx"
