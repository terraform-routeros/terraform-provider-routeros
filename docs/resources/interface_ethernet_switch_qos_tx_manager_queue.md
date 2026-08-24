# routeros_interface_ethernet_switch_qos_tx_manager_queue (Resource)

Configures the egress scheduler for one traffic class within a tx-manager. The
eight queues (traffic class 0..7) are auto-created by the switch chip. This
resource **adopts** an existing queue keyed by `traffic_class` (within
`tx_manager`) and can only modify it — it is never added or deleted.

## Example Usage
```terraform
resource "routeros_interface_ethernet_switch_qos_tx_manager_queue" "tc6" {
  traffic_class = 6
  schedule      = "high-priority-group"
  weight        = 16
}
```

## Schema

### Required

- `traffic_class` (Number) Traffic class (egress queue index, 0..7) this queue configures.

### Optional

- `default` (Boolean) Whether this queue is the default queue.
- `ecn` (Boolean) Enable Explicit Congestion Notification (ECN) marking for this queue.
- `queue_buffers` (String) Amount of buffers dedicated to this queue (`auto` or a numeric value).
- `schedule` (String) Scheduling group for this queue. One of `high-priority-group`, `low-priority-group`, `strict-priority`.
- `shared_pool_index` (Number) Index of the shared buffer pool used when `use_shared_buffers` is enabled.
- `tx_manager` (String) Name of the tx-manager (egress scheduler) this queue belongs to. Default `default`.
- `use_shared_buffers` (Boolean) Whether this queue may use the shared buffer pool.
- `weight` (Number) Weight of this queue within its scheduling group (ignored for `strict-priority`).
- `wred` (Boolean) Enable Weighted Random Early Detection (WRED) for this queue.

### Read-Only

- `id` (String) The ID of this resource.

## Import
```shell
terraform import routeros_interface_ethernet_switch_qos_tx_manager_queue.tc6 "traffic-class=6"
```
