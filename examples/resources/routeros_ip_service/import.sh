# Import with the name of the ip service in case of the example use www-ssl
terraform import routeros_ip_service.www_ssl www-ssl#Or you can import a resource using one of its attributes
terraform import routeros_ip_service.www_ssl "name=xxx"
