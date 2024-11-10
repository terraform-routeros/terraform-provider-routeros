# routeros_ip_ipsec_mode_config (Resource)


## Example Usage
```terraform
resource "routeros_ip_ipsec_mode_config" "test" {
  name          = "test-cfg"
  address       = "1.2.3.4"
  split_include = ["0.0.0.0/0"]
  split_dns     = ["1.1.1.1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `address` (String) Single IP address for the initiator instead of specifying a whole address pool.
- `address_pool` (String) Name of the address pool from which the responder will try to assign address if mode-config is enabled.
- `address_prefix_length` (Number) Prefix length (netmask) of the assigned address from the pool.
- `connection_mark` (String) Firewall connection mark.
- `responder` (Boolean) Specifies whether the configuration will work as an initiator (client) or responder (server). The initiator will request for mode-config parameters from the responder.
- `split_dns` (Set of String) List of DNS names that will be resolved using a system-dns=yes or static-dns= setting.
- `split_include` (Set of String) List of subnets in CIDR format, which to tunnel. Subnets will be sent to the peer using the CISCO UNITY extension, a remote peer will create specific dynamic policies.
- `src_address_list` (String) Specifying an address list will generate dynamic source NAT rules. This parameter is only available with responder=no. A roadWarrior client with NAT.
- `static_dns` (String) Manually specified DNS server's IP address to be sent to the client.
- `system_dns` (Boolean) When this option is enabled DNS addresses will be taken from `/ip dns`.
- `use_responder_dns` (String)

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/ipsec/mode/config get [print show-ids]]
terraform import routeros_ip_ipsec_mode_config.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ip_ipsec_mode_config.test "address=1.2.3.4"
```