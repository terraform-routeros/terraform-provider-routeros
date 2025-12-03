#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/user-manager/router get [print show-ids]]
terraform import routeros_user_manager_router.test '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_user_manager_router.test "name=xxx"
