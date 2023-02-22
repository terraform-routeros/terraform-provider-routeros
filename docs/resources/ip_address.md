---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "routeros_ip_address Resource - terraform-provider-routeros"
subcategory: "IP"
description: |-
  
---

# routeros_ip_address (Resource)


```terraform
resource "routeros_ip_address" "address" {
  address   = "10.0.0.1/24"
  interface = "bridge"
  network   = "10.0.0.0"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String)
- `interface` (String)
- `network` (String)

### Optional

- `comment` (String)
- `disabled` (Boolean)

### Read-Only

- `actual_interface` (String)
- `id` (String) The ID of this resource.
- `invalid` (Boolean)

