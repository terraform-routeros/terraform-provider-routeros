#The ID can be found via API or the terminal
#The command for the terminal is -> /routing/filter/rule/print show-ids
terraform import routeros_routing_filter_rule.test "*0"#Or you can import a resource using one of its attributes
terraform import routeros_routing_filter_rule.test "name=xxx"
