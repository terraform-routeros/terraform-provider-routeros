# routeros_interface_bridge_settings (Resource)


## Example Usage
```terraform
resource "routeros_interface_bridge_settings" "settings" {
  use_ip_firewall = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `allow_fast_path` (Boolean) Whether to enable a bridge FastPath globally.
- `use_ip_firewall` (Boolean) Force bridged traffic to also be processed by prerouting, forward and postrouting sections of IP routing ( Packet Flow). This does not apply to routed traffic. This property is required in case you want to assign Simple Queues or global Queue Tree to traffic in a bridge. Property use-ip-firewall-for-vlan is required in case bridge vlan-filtering is used.
- `use_ip_firewall_for_pppoe` (Boolean) Send bridged un-encrypted PPPoE traffic to also be processed by IP/Firewall. This property only has effect when use-ip-firewall is set to yes. This property is required in case you want to assign Simple Queues or global Queue Tree to PPPoE traffic in a bridge.
- `use_ip_firewall_for_vlan` (Boolean) Send bridged VLAN traffic to also be processed by IP/Firewall. This property only has effect when use-ip-firewall is set to yes. This property is required in case you want to assign Simple Queues or global Queue Tree to VLAN traffic in a bridge.

### Read-Only

- `bridge_fast_forward_bytes` (Number) Shows byte count forwarded by Bridge Fast Forward.
- `bridge_fast_forward_packets` (Number) Shows packet count forwarded by Bridge Fast Forward.
- `bridge_fast_path_active` (Boolean) Shows whether a bridge FastPath is active globally, FastPatch status per bridge interface is not displayed.
- `bridge_fast_path_bytes` (Number) Shows byte count forwarded by Bridge Fast Path.
- `bridge_fast_path_packets` (Number) Shows packet count forwarded by Bridge FastPath.
- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
terraform import routeros_interface_bridge_settings.settings .
```
