#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/user-manager/attribute get [print show-ids]]
terraform routeros_user_manager_attribute.mikrotik_wireless_comment '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_user_manager_attribute.mikrotik_wireless_comment "name=xxx"
