# routeros_interface_ethernet_switch_qos_port (Resource)

Configures per-port QoS on the switch chip: which default profile to use and
whether to trust the incoming L2 (PCP) and/or L3 (DSCP) markings.

Port rows are auto-created by the switch chip (one per port). This resource
**adopts** an existing row keyed by `name` and can only modify it — it is never
added or deleted. Removing the resource from Terraform leaves the row in place.

## Example Usage
```terraform
resource "routeros_interface_ethernet_switch_qos_port" "sfp1" {
  name     = "sfp-sfpplus1"
  profile  = "5-Standard (BE/Default)"
  trust_l2 = "trust"
  trust_l3 = "trust"
}
```

## Schema

### Required

- `name` (String) Name of the switch port to configure.

### Optional

- `map` (String) Name of the QoS map set used for classification on this port.
- `pfc` (String) Priority Flow Control (PFC) mode for this port.
- `profile` (String) Default `qos profile` applied to traffic on this port.
- `trust_l2` (String) Whether to trust the L2 (802.1p PCP) priority of ingress traffic on this port. One of `trust`, `ignore`, `keep`.
- `trust_l3` (String) Whether to trust the L3 (IP DSCP) priority of ingress traffic on this port. One of `trust`, `ignore`, `keep`.
- `tx_manager` (String) Name of the `qos tx-manager` (egress scheduler) applied to this port.

### Read-Only

- `id` (String) The ID of this resource.
- `invalid` (Boolean)
- `running` (Boolean)
- `switch` (String) Name of the switch this port belongs to.

## Import
```shell
terraform import routeros_interface_ethernet_switch_qos_port.sfp1 "name=sfp-sfpplus1"
```
