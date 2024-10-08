# routeros_capsman_manager_interface (Resource)


## Example Usage
```terraform
resource "routeros_capsman_manager_interface" "test_manager_interface" {
  interface = "ether1"
  forbid    = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface` (String) Name of the interface.

### Optional

- `comment` (String)
- `disabled` (Boolean)
- `forbid` (Boolean) Disable interface listening.

### Read-Only

- `default` (Boolean) It's the default item.
- `dynamic` (Boolean) Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is ->  :put [/caps-man/manager/interface get [print show-ids]]
terraform import routeros_capsman_manager_interface.test_manager_interface "*6"
```
