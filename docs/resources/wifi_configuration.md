# routeros_wifi_configuration (Resource)
*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*

## Example Usage
```terraform
resource "routeros_wifi_aaa" "aaa1" {
  called_format   = "S"
  name            = "aaa1"
  password_format = ""
  username_format = "AA:AA:AA:AA:AA:AA"
}

resource "routeros_wifi_channel" "channel1" {
  name                = "1"
  band                = "2ghz-n"
  frequency           = [2412]
  secondary_frequency = ["disabled"]
  skip_dfs_channels   = "disabled"
  width               = "20mhz"
}

resource "routeros_wifi_datapath" "datapath1" {
  name             = "datapath1"
  bridge           = "bridge1"
  client_isolation = false
}

resource "routeros_wifi_security" "security1" {
  name                 = "security1"
  authentication_types = ["wpa2-psk", "wpa3-psk"]
  ft                   = true
  ft_preserve_vlanid   = true
  passphrase           = "password"
}

resource "routeros_wifi_configuration" "configuration1" {
  country = "Netherlands"
  manager = "capsman"
  mode    = "ap"
  name    = "configuration1"
  ssid    = "my-network"

  aaa = {
    config = routeros_wifi_aaa.aaa1.name
  }
  channel = {
    config = routeros_wifi_channel.channel1.name
  }
  datapath = {
    config = routeros_wifi_datapath.datapath1.name
  }
  security = {
    config = routeros_wifi_security.security1.name
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the configuration.

### Optional

- `aaa` (Map of String) AAA inline settings.
- `antenna_gain` (Number) An option overrides the default antenna gain.
- `beacon_interval` (String) Time interval between beacon frames.
- `chains` (Set of Number) Radio chains to use for receiving signals.
- `channel` (Map of String) Channel inline settings.
- `comment` (String)
- `country` (String) An option determines which regulatory domain restrictions are applied to an interface.
- `datapath` (Map of String) Datapath inline settings.
- `disabled` (Boolean)
- `dtim_period` (Number) A period at which to transmit multicast traffic, when there are client devices in power save mode connected to the AP.
- `hide_ssid` (Boolean) This property has effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.
- `interworking` (Map of String) Interworking inline settings.
- `manager` (String) An option to specify the remote CAP mode.
- `mode` (String) An option to specify the access point operational mode.
- `multicast_enhance` (String) An option to enable converting every multicast-address IP or IPv6 packet into multiple unicast-addresses frames for each connected station.
- `qos_classifier` (String) An option to specify the QoS classifier.
- `security` (Map of String) Security inline settings.
- `ssid` (String) SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network.
- `steering` (Map of String) Steering inline settings.
- `tx_chains` (Set of Number) Radio chains to use for transmitting signals.
- `tx_power` (Number) A limit on the transmit power (in dBm) of the interface.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wifi/configuration get [print show-ids]]
terraform import routeros_wifi_configuration.configuration1 '*1'
```
