# routeros_interface_ethernet_switch_qos_map_ip (Resource)

Maps an ingress DSCP value to a QoS profile within a named map (usually
`default`). This is how L3 (IP) traffic classification is programmed into the
switch chip.

Requires `qos-hw-offloading=yes` on the switch chip
(`routeros_interface_ethernet_switch`).

## Example Usage
```terraform
resource "routeros_interface_ethernet_switch_qos_map_ip" "ef" {
  dscp    = 46
  profile = "1-Real-Time (EF/VoIP/Video)"
}
```

## Schema

### Required

- `profile` (String) Name of the `qos profile` this DSCP value is mapped to.

### Optional

- `disabled` (Boolean)
- `dscp` (Number) Ingress Differentiated Services Code Point (DSCP) value to match.
- `map` (String) Name of the QoS map set this entry belongs to. Default `default`.

### Read-Only

- `id` (String) The ID of this resource.

## Import
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ethernet/switch/qos/map/ip get [print show-ids]]
terraform import routeros_interface_ethernet_switch_qos_map_ip.ef "*0"
#or by field
terraform import routeros_interface_ethernet_switch_qos_map_ip.ef "dscp=46"
```
