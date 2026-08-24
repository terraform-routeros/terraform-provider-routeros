# routeros_interface_ethernet_switch_qos_profile (Resource)

Maps a DSCP/PCP value to an internal switch-chip traffic class (egress queue).
Profiles are the building block referenced by `routeros_interface_ethernet_switch_qos_map_ip`
and by the per-port `routeros_interface_ethernet_switch_qos_port` configuration.

Requires `qos-hw-offloading=yes` on the switch chip
(`routeros_interface_ethernet_switch`). Available on CRS3xx/CRS5xx and similar
switches.

## Example Usage
```terraform
resource "routeros_interface_ethernet_switch_qos_profile" "realtime" {
  name          = "1-Real-Time (EF/VoIP/Video)"
  dscp          = 46
  traffic_class = 7
}
```

## Schema

### Required

- `name` (String) Profile name, referenced from the `qos map ip` and `qos port` menus.

### Optional

- `disabled` (Boolean)
- `dscp` (Number) Differentiated Services Code Point (DSCP) value assigned to this profile.
- `pcp` (Number) Priority Code Point (802.1p) value assigned to this profile.
- `traffic_class` (Number) Internal switch-chip traffic class (egress queue) that traffic matching this profile is mapped to.

### Read-Only

- `id` (String) The ID of this resource.

## Import
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/ethernet/switch/qos/profile get [print show-ids]]
terraform import routeros_interface_ethernet_switch_qos_profile.realtime "*1"
#or by name
terraform import routeros_interface_ethernet_switch_qos_profile.realtime "name=1-Real-Time (EF/VoIP/Video)"
```
