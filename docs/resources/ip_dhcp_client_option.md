# routeros_ip_dhcp_client_option (Resource)


## Example Usage
```terraform
resource "routeros_ip_dhcp_client_option" "option" {
  name = "my-dhcp-option"
  code = 60
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (Number) The dhcp-client option code.
- `name` (String) The name that will be used in dhcp-client.

### Optional

- `raw_value` (String) raw_value is computed from value.
- `value` (String) The dhcp-client option

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-client/option get [print show-ids]]
terraform import routeros_ip_dhcp_client_option.option "*0"
```
