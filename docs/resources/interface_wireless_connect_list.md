# routeros_interface_wireless_connect_list (Resource)


## Example Usage
```terraform
resource "routeros_interface_wireless_connect_list" "test" {
  interface        = "wlan0"
  security_profile = "test-secp"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface` (String) Name of the interface.

### Optional

- `allow_signal_out_of_range` (String)
- `area_prefix` (String) Rule matches if area value of AP (a proprietary extension) begins with specified value.area value is a proprietary extension.
- `comment` (String)
- `connect` (Boolean) Available options: yes - Connect to access point that matches this rule. no - Do not connect to any access point that matches this rule.
- `disabled` (Boolean)
- `interworking` (String)
- `iw_asra` (String) Additional Steps Required for Access. Set to yes, if a user should take additional steps to access the internet, like the walled garden.
- `iw_authentication_types` (String) This property is only effective when `asra` is set to `yes`. Value of `url` is optional and not needed if `dns-redirection` or `online-enrollment` is selected. To set the value of `url` to empty string use double quotes. For example: `authentication-types=online-enrollment:""`
- `iw_connection_capabilities` (String) This option allows to provide information about the allowed IP protocols and ports. This information can be provided in ANQP response. The first number represents the IP protocol number, the second number represents a port number.
  * closed - set if protocol and port combination is not allowed;
  * open - set if protocol and port combination is allowed;
  * unknown - set if protocol and port combination is either open or closed.

Example: `connection-capabilities=6:80:open,17:5060:closed`Setting such a value on an Access Point informs the Wireless client, which is connecting to the Access Point, that HTTP (6 - TCP, 80 - HTTP) is allowed and VoIP (17 - UDP; 5060 - VoIP) is not allowed. This property does not restrict or allow usage of these protocols and ports, it only gives information to station device which is connecting to Access Point.
- `iw_esr` (String) Emergency services reachable (ESR). Set to yes in order to indicate that emergency services are reachable through the access point.
- `iw_hessid` (String) Homogenous extended service set identifier (HESSID). Devices that provide access to same external networks are in one homogenous extended service set. This service set can be identified by HESSID that is the same on all access points in this set. 6-byte value of HESSID is represented as MAC address. It should be globally unique, therefore it is advised to use one of the MAC address of access point in the service set.
- `iw_hotspot20` (String) Indicate Hotspot 2.0 capability of the Access Point.
- `iw_hotspot20_dgaf` (String) Downstream Group-Addressed Forwarding (DGAF). Sets value of DGAF bit to indicate whether multicast and broadcast frames to clients are disabled or enabled.
  * yes - multicast and broadcast frames to clients are enabled;
  * no - multicast and broadcast frames to clients are disabled.
To disable multicast and broadcast frames set `multicast-helper=full`.
- `iw_internet` (String) Whether the internet is available through this connection or not. This information is included in the Interworking element.
- `iw_ipv4_availability` (String) Information about what IPv4 address and access are available.
  * not-available - Address type not available;
  * public - public IPv4 address available;
  * port-restricted - port-restricted IPv4 address available;
  * single-nated - single NATed private IPv4 address available;
  * double-nated - double NATed private IPv4 address available;
  * port-restricted-single-nated -port-restricted IPv4 address and single NATed IPv4 address available;
  * port-restricted-double-nated - port-restricted IPv4 address and double NATed IPv4 address available;
  * unknown - availability of the address type is not known.
- `iw_ipv6_availability` (String) Information about what IPv6 address and access are available.
  * not-available - Address type not available;
  * available - address type available;
  * unknown - availability of the address type is not known.
- `iw_network_type` (String) Information about network access type.
  * emergency-only - a network dedicated and limited to accessing emergency services;
  * personal-device - a network of personal devices. An example of this type of network is a camera that is attached to a printer, thereby forming a network for the purpose of printing pictures;
  * private - network for users with user accounts. Usually used in enterprises for employees, not guests;
  * private-with-guest - same as private, but guest accounts are available;
  * public-chargeable - a network that is available to anyone willing to pay. For example, a subscription to Hotspot 2.0 service or in-room internet access in a hotel;
  * public-free - network is available to anyone without any fee. For example, municipal network in city or airport Hotspot;
  * test - network used for testing and experimental uses. Not used in production;
  * wildcard - is used on Wireless clients. Sending probe request with a wildcard as network type value will make all Interworking Access Points respond despite their actual network-type setting.

A client sends a probe request frame with network-type set to value it is interested in. It will receive replies only from access points with the same value (except the case of wildcard).
- `iw_realms` (String) Information about supported realms and the corresponding EAP method. `realms=example.com:eap-tls,foo.ba:not-specified`
- `iw_roaming_ois` (String) Organization identifier (OI) usually are 24-bit is unique identifiers like organizationally unique identifier (OUI) or company identifier (CID). In some cases, OI is longer for example OUI-36.A subscription service provider (SSP) can be specified by its OI. roaming-ois property can contain zero or more SSPs OIs whose networks are accessible via this AP. Length of OI should be specified before OI itself. For example, to set E4-8D-8C and 6C-3B-6B: `roaming-ois=03E48D8C036C3B6B`
- `iw_uesa` (String) Unauthenticated emergency service accessible (UESA).
  * no - indicates that no unauthenticated emergency services are reachable through this Access Point;
  * yes - indicates that higher layer unauthenticated emergency services are reachable through this Access Point.
- `iw_venue` (String) Specify the venue in which the Access Point is located. Choose the value from available ones. Some examples:
   ```
  venue=business-bank
  venue=mercantile-shopping-mall
  venue=educational-university-or-college
   ```
- `mac_address` (String) Rule matches only AP with the specified MAC address.
- `security_profile` (String) Name of security profile that is used when connecting to matching access points, If value of this property is none, then security profile specified in the interface configuration will be used. In station mode, rule will match only access points that can support specified security profile. Value none will match access point that supports security profile that is specified in the interface configuration. In access point mode value of this property will not be used to match remote devices.
- `signal_range` (String) Rule matches if signal strength of the access point is within the range. If station establishes connection to access point that is matched by this rule, it will disconnect from that access point when signal strength goes out of the specified range.
- `ssid` (String) Rule matches access points that have this SSID. Empty value matches any SSID. This property has effect only when station mode interface ssid is empty, or when access point mode interface has wds-ignore-ssid=yes.
- `three_gpp` (String)
- `wireless_protocol` (String)

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wireless/connect-list get [print show-ids]]
terraform import routeros_interface_wireless_connect_list.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_interface_wireless_connect_list.test "name=xxx"
```