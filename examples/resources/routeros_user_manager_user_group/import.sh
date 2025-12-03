#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/user-manager/user/group get [print show-ids]]
terraform import routeros_user_manager_user_group.test '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_user_manager_user_group.test "name=xxx"
