#The ID can be found via API or the terminal
#The command for the terminal is ->  :put [/caps-man/provisioning get [print show-ids]]
terraform import routeros_capsman_provisioning.test_provisioning "*B"#Or you can import a resource using one of its attributes
terraform import routeros_capsman_provisioning.test_provisioning "name=xxx"
