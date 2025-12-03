#The ID can be found via API or the terminal
#The command for the terminal is -> /routing/ospf/instance/print show-ids
terraform import routeros_routing_ospf_instance.test_routing_ospf_instance "*0"
#Or you can import a resource using one of its attributes
terraform import routeros_routing_ospf_instance.test_routing_ospf_instance "name=xxx"
