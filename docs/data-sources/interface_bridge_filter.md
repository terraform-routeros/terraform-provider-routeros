# routeros_interface_bridge_filter (Data Source)


## Example Usage
```terraform
data "routeros_ip_firewall_filter" "rules" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Map of String) Additional request filtering options.

### Read-Only

- `filters` (List of Object) (see [below for nested schema](#nestedatt--filters))
- `id` (String) The ID of this resource.

<a id="nestedatt--filters"></a>
### Nested Schema for `filters`

Read-Only:

- `action` (String)
- `bytes` (Number)
- `chain` (String)
- `comment` (String)
- `dynamic` (Boolean)
- `id` (String)
- `in_interface` (String)
- `invalid` (Boolean)
- `mac_protocol` (String)
- `packets` (Number)

