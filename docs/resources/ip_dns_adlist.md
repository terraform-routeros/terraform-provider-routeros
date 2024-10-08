# routeros_ip_dns_adlist (Resource)


## Example Usage
```terraform
resource "routeros_ip_dns_adlist" "test" {
  url        = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
  ssl_verify = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disabled` (Boolean)
- `file` (String) Used to specify a local file path from which to read adlist data.
- `ssl_verify` (Boolean) Specifies whether to validate the server's SSL certificate when connecting to an online resource. Will use the `/certificate` list to verify server validity.
- `url` (String) Used to specify the URL of an adlist.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dns/adlist get [print show-ids]]
terraform import routeros_ip_dns_adlist.test "*0"
```
