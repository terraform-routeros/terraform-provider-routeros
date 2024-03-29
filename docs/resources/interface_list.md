# routeros_interface_list (Resource)


## Example Usage
```terraform
resource "routeros_interface_list" "list" {
  name = "my-list"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Changing the name of this resource will force it to be recreated.
	> The links of other configuration properties to this resource may be lost!
	> Changing the name of the resource outside of a Terraform will result in a loss of control integrity for that resource!

### Optional

- `comment` (String)
- `exclude` (String)
- `include` (String)

### Read-Only

- `builtin` (Boolean)
- `dynamic` (Boolean) Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/list get [print show-ids]]
terraform import routeros_interface_list.list "*2000010"
```
