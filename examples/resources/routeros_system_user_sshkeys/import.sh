#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/user/ssh-keys get [print show-ids]]
terraform import routeros_system_user_sshkeys.test *1#Or you can import a resource using one of its attributes
terraform import routeros_system_user_sshkeys.test "name=xxx"
