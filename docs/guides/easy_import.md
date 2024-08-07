# Install package

Original [issue](https://github.com/terraform-routeros/terraform-provider-routeros/issues/488)

## Example
```shell
#!/bin/bash

USER=admin
PASS=
HOST=http://router.local

i=0
curl -s -u ${USER}:${PASS} ${HOST}/rest/ip/firewall/address-list | jq -c '.[] | select(.dynamic | ascii_downcase == "false") | {index: .".id", address: .address, comment: .comment, list: .list}'  | while read rec; do
  index=$(echo $rec | jq .index)
  idx=$(printf "%00004d" $i)
  # echo $rec
  bash -cv "tofu state rm 'module.dev-gw0.routeros_ip_firewall_addr_list.address_list[\"$idx\"]'"
  bash -cv "tofu import 'module.dev-gw0.routeros_ip_firewall_addr_list.address_list[\"$idx\"]' $index"
  let i=${i}+1
done
```

```terraform
variable "address_list" {
  type = list(object({
    address = string
    comment = optional(string)
    disabled = optional(bool, false)
    dynamic  = optional(bool, false)
    list     = string
  }))

  default = [
    { address="192.168.88.11", comment="example 2", list="srv" },
    { address="192.168.88.12", comment="example 2", list="srv" },
    { address="192.168.88.1", comment="example", list="routeros" },
]

locals {
  # https://discuss.hashicorp.com/t/does-map-sort-keys/12056/2
  # Map keys are always iterated in lexicographical order!
  address_list_map = { for idx, rule in var.address_list : format("%00004d", idx) => rule }
}

resource "routeros_ip_firewall_addr_list" "address_list" {
  for_each = local.address_list_map
  address  = each.value.address
  comment  = each.value.comment
  disabled = each.value.disabled
  list     = each.value.list
}
```