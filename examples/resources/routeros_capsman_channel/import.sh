#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/caps-man/channel get [print show-ids]]
terraform import routeros_capsman_channel.test_channel "*1"#Or you can import a resource using one of its attributes
terraform import routeros_capsman_channel.test_channel "name=xxx"
