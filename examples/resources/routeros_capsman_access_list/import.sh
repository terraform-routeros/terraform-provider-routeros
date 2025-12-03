#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/caps-man/access-list get [print show-ids]]
terraform import routeros_capsman_access_list.test_rule "*1"
#Or you can import a resource using one of its attributes
terraform import routeros_capsman_access_list.test_rule "name=xxx"
