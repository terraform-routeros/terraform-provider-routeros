#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/zerotier/controller get [print show-ids]]
terraform import routeros_zerotier_controller.test '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_zerotier_controller.test "name=xxx"
