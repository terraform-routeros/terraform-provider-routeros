# routeros_capsman_interface (Resource)


## Example Usage
```terraform
resource "routeros_capsman_channel" "channel1" {
  name      = "1"
  band      = "2ghz-g/n"
  frequency = [2412]
}

resource "routeros_capsman_interface" "cap1" {
  name = "cap1"

  channel = {
    config = routeros_capsman_channel.channel1.name
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the interface.

### Optional

- `arp_timeout` (String) ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in IP/Settings, default is 30s. Can use postfix `ms`, `s`, `m`, `h`, `d` for milliseconds, seconds, minutes, hours or days. If no postfix is set then seconds (s) is used.
- `channel` (Map of String) Channel inline settings.
- `comment` (String)
- `configuration` (Map of String) Configuration inline settings.
- `datapath` (Map of String) Datapath inline settings.
- `disabled` (Boolean)
- `mac_address` (String) MAC address (BSSID) to use for the interface.
- `master_interface` (String) The corresponding master interface of the virtual one.
- `radio_mac` (String) The MAC address of the associated radio.
- `radio_name` (String) Name of the associated radio.
- `rates` (Map of String) Rates inline settings.
- `security` (Map of String) Security inline settings.

### Read-Only

- `bound` (Boolean) A flag whether the interface is currently available for the CAPsMAN.
- `id` (String) The ID of this resource.
- `inactive` (Boolean) A flag whether the interface is currently inactive.
- `l2mtu` (Number) Layer2 Maximum transmission unit. [See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).
- `master` (Boolean) A flag whether the interface is not a virtual one.
- `running` (Boolean) A flag whether the interface has established a link to another device.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/caps-man/interface get [print show-ids]]
terraform import routeros_capsman_interface.cap1 '*1'
```
