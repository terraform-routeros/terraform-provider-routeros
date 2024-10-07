resource "routeros_ip_ipsec_policy_group" "group-for-policy" {
  name = "test-group"
}

resource "routeros_ip_ipsec_policy" "policy" {
  dst_address = "0.0.0.0/0"
  group       = routeros_ip_ipsec_policy_group.group-for-policy.name
  proposal    = "NordVPN"
  src_address = "0.0.0.0/0"
  template    = true
}
