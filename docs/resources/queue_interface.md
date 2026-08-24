# routeros_queue_interface (Resource)

Assigns an interface queue type to an interface (`/queue interface`). The rows
are auto-created by RouterOS (one per interface). This resource **adopts** an
existing row keyed by `interface` and can only modify it — it is never added or
deleted.

## Example Usage
```terraform
resource "routeros_queue_interface" "ether1" {
  interface = "ether1"
  queue     = "ethernet-default"
}
```

## Schema

### Required

- `interface` (String) Name of the interface whose queue is being configured.
- `queue` (String) Name of the interface queue type to assign (e.g. `ethernet-default`, `only-hardware-queue`, `no-queue`).

### Read-Only

- `active_queue` (String) Queue type currently active on the interface.
- `default_queue` (String) Default queue type for the interface.
- `id` (String) The ID of this resource.

## Import
```shell
terraform import routeros_queue_interface.ether1 "interface=ether1"
```
