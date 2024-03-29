# routeros_system_user_aaa (Resource)


## Example Usage
```terraform
resource "routeros_system_user_aaa" "settings" {
  use_radius = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `accounting` (Boolean) An option that enables accounting for users.
- `default_group` (String) The user group that is used by default for users authenticated via a RADIUS server.
- `exclude_groups` (Set of String) A set of groups that are not allowed for users authenticated by RADIUS.
- `interim_update` (String) Interval between scheduled RADIUS Interim-Update messages.
- `use_radius` (Boolean) An option whether to use RADIUS server.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
terraform import routeros_system_user_aaa.settings .
```
