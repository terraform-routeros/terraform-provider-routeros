# routeros_ip_hotspot_service_port (Resource)


## Example Usage
```terraform
resource "routeros_ip_hotspot_service_port" "test" {
  name     = "ftp"
  disabled = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Service name.

### Optional

- `disabled` (Boolean)

### Read-Only

- `id` (String) The ID of this resource.
- `ports` (String)

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/hotspot/service-port get [print show-ids]]
terraform import routeros_ip_hotspot_service_port.test *1
```