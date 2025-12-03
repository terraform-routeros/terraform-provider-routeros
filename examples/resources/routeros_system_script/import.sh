#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/system/script get [print show-ids]]
terraform import routeros_system_script.script "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_system_script.script "name=xxx"
