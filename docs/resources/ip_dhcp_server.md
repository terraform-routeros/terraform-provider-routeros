---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "routeros_ip_dhcp_server Resource - terraform-provider-routeros"
subcategory: "IP"
description: |-
  
---

# routeros_ip_dhcp_server (Resource)


```terraform
resource "routeros_ip_dhcp_server" "server" {
  address_pool = "my_address_pool"
  interface    = "bridge"
  name         = "bridge_dhcp"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address_pool` (String)
- `interface` (String)
- `name` (String)

### Optional

- `authoritative` (Boolean)
- `disabled` (Boolean)
- `lease_script` (String)
- `lease_time` (String)
- `use_radius` (Boolean)

### Read-Only

- `dynamic` (Boolean)
- `id` (String) The ID of this resource.
- `invalid` (Boolean)

