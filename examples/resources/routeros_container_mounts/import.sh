# Import with the name of the container mount in case of the example use Caddyfile
terraform import routeros_container_mounts.caddyfile Caddyfile
#Or you can import a resource using one of its attributes
terraform import routeros_container_mounts.caddyfile "name=xxx"
