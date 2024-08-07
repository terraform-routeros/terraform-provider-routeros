resource "routeros_capsman_configuration" "test_configuration" {
  comment              = "Comment"
  country              = "no_country_set"
  disconnect_timeout   = "1s150ms"
  distance             = "indoors"
  frame_lifetime       = "0.12" // 120ms
  guard_interval       = "long"
  hide_ssid            = true
  hw_protection_mode   = "rts-cts"
  hw_retries           = 1
  installation         = "indoor"
  keepalive_frames     = "enabled"
  load_balancing_group = ""
  max_sta_count        = 1
  mode                 = "ap"
  multicast_helper     = "full"
  name                 = "test_configuration"
  rx_chains            = [1, 3]
  ssid                 = "SSID"
  tx_chains            = [0, 2]
}

resource "routeros_capsman_channel" "test_channel" {
  name = "test-channel-config"
}

resource "routeros_capsman_datapath" "test_datapath" {
  name = "test-datapath-config"
}

resource "routeros_capsman_rates" "test_rates" {
  name = "test-rates-config"
}

resource "routeros_capsman_security" "test_security" {
  name = "test-security-config"
}

resource "routeros_capsman_configuration" "test_configuration_2" {
  name = "test_configuration_name"

  channel = {
    config                = "${routeros_capsman_channel.test_channel.name}"
    band                  = "2ghz-b/g/n"
    control_channel_width = "10mhz"
    extension_channel     = "eCee"
    frequency             = 2412
    reselect_interval     = "1h"
    save_selected         = "true"
    secondary_frequency   = "disabled"
    skip_dfs_channels     = "true"
    tx_power              = 20
  }

  datapath = {
    config                      = "${routeros_capsman_datapath.test_datapath.name}"
    arp                         = "local-proxy-arp"
    bridge                      = "bridge"
    bridge_cost                 = "100"
    bridge_horizon              = "200"
    client_to_client_forwarding = "true"
    interface_list              = "static"
    l2mtu                       = "1450"
    local_forwarding            = "true"
    mtu                         = "1500"
    vlan_id                     = "101"
    vlan_mode                   = "no-tag"
    //   openflow_switch             = "aaa"
  }

  rates = {
    config            = "${routeros_capsman_rates.test_rates.name}"
    basic             = "1Mbps,5.5Mbps,6Mbps,18Mbps,36Mbps,54Mbps"
    ht_basic_mcs      = "mcs-0,mcs-7,mcs-11,mcs-14,mcs-16,mcs-21"
    ht_supported_mcs  = "mcs-3,mcs-8,mcs-10,mcs-13,mcs-17,mcs-18"
    supported         = "2Mbps,11Mbps,9Mbps,12Mbps,24Mbps,48Mbps"
    vht_basic_mcs     = "none"
    vht_supported_mcs = "mcs0-9,mcs0-7"
  }

  security = {
    config                = "${routeros_capsman_security.test_security.name}"
    authentication_types  = "wpa-psk,wpa-eap"
    disable_pmkid         = "true"
    eap_methods           = "eap-tls,passthrough"
    eap_radius_accounting = "true"
    encryption            = "aes-ccm,tkip"
    group_encryption      = "aes-ccm"
    group_key_update      = "1h"
    passphrase            = "AAAAAAAAA"
    tls_certificate       = "none"
    tls_mode              = "verify-certificate"
  }

  depends_on = [
    routeros_capsman_channel.test_channel,
    routeros_capsman_datapath.test_datapath,
    routeros_capsman_rates.test_rates,
    routeros_capsman_security.test_security
  ]
}
