#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/container/envs get [print show-ids]]
terraform import routeros_container_envs.test_envs "*1"
#Or you can import a resource using one of its attributes
terraform import routeros_container_envs.test_envs "name=xxx"
