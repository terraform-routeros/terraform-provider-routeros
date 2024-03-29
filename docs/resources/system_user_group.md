# routeros_system_user_group (Resource)


## Example Usage
```terraform
resource "routeros_system_user_group" "terraform" {
  name   = "terraform"
  policy = ["api", "!ftp", "!local", "password", "policy", "read", "!reboot", "!rest-api", "!romon", "sensitive", "!sniff", "!ssh", "!telnet", "!test", "!web", "!winbox", "write"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the user group

### Optional

- `comment` (String)
- `policy` (Set of String) A set of allowed policies.
- `skin` (String) The name of the skin that will be used for WebFig.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/user/group get [print show-ids]]
terraform import routeros_system_user_group.terraform *1
```
