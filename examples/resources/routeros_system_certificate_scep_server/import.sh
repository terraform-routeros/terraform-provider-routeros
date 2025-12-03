# The ID can be found via API or the terminal
# The command for the terminal is -> /certificate/scep-server/print show-ids
terraform import routeros_system_certificate_scep_server.example_scep_server "*1"
#Or you can import a resource using one of its attributes
terraform import routeros_system_certificate_scep_server.example_scep_server "name=xxx"
