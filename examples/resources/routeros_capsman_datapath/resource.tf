resource "routeros_capsman_datapath" "test_datapath" {
  name                        = "test_datapath"
  comment                     = "test_datapath"
  arp                         = "local-proxy-arp"
  bridge                      = "bridge"
  bridge_cost                 = 100
  bridge_horizon              = 200
  client_to_client_forwarding = true
  interface_list              = "static"
  l2mtu                       = 1450
  local_forwarding            = true
  mtu                         = 1500
  vlan_id                     = 101
  vlan_mode                   = "no-tag"
  //  openflow_switch             = "aaa"
}