#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/queue/tree get [print show-ids]]
terraform import routeros_queue_tree.test *1000000
#Or you can import a resource using one of its attributes
terraform import routeros_queue_tree.test "name=server"
