#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/caps-man/rates get [print show-ids]]
terraform import routeros_capsman_rates.test_rates "*1"#Or you can import a resource using one of its attributes
terraform import routeros_capsman_rates.test_rates "name=xxx"
