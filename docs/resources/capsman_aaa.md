# routeros_capsman_aaa (Resource)


## Example Usage
```terraform
resource "routeros_capsman_aaa" "test_3a" {
  called_format = "ssid"
  mac_mode      = "as-username-and-password"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `called_format` (String) Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius server clients, you can specify 'called-id' in order to separate multiple entires.
- `interim_update` (String) When RADIUS accounting is used, Access Point periodically sends accounting information updates to the RADIUS server. This property specifies the default update interval that can be overridden by the RADIUS server using the Acct-Interim-Interval attribute.
- `mac_caching` (String) If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication responses for a specified time, and will not contact the RADIUS server if matching cache entry already exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.
- `mac_format` (String) Controls how the MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC authentication and MAC accounting RADIUS requests.
- `mac_mode` (String) By default Access Point uses an empty password, when sending Access-Request during MAC authentication. When this property is set to as-username-and-password, Access Point will use the same value for the User-Password attribute as for the User-Name attribute.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
terraform import routeros_capsman_aaa.test_3a .
```
