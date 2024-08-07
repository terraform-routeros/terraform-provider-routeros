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
