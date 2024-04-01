resource "routeros_ip_neighbor_discovery_settings" "test" {
  discover_interface_list  = "static"
  lldp_med_net_policy_vlan = "1"
  mode                     = "tx-and-rx"
  protocol                 = []
}
