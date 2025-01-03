# If you need to add a reference to an existing configuration, each inline section contains a `config` parameter 
# where you can specify the name of the actual resource.
# configuration = {
#   config = routeros_wifi_configuration.my-config.name
# }
resource "routeros_wifi" "wifi1" {
  configuration = {
    manager = "capsman"
  }
  name = "wifi1"
}
