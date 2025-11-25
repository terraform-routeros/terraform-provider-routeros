package main

var (
	providerAliasesMap = map[string]struct{}{
		"routeros_dhcp_client":         struct{}{},
		"routeros_dhcp_client_option":  struct{}{},
		"routeros_dhcp_server":         struct{}{},
		"routeros_dhcp_server_network": struct{}{},
		"routeros_dhcp_server_lease":   struct{}{},
		"routeros_firewall_addr_list":  struct{}{},
		"routeros_firewall_filter":     struct{}{},
		"routeros_firewall_mangle":     struct{}{},
		"routeros_firewall_nat":        struct{}{},
		"routeros_dns":                 struct{}{},
		"routeros_dns_record":          struct{}{},
		"routeros_bridge":              struct{}{},
		"routeros_bridge_mlag":         struct{}{},
		"routeros_bridge_port":         struct{}{},
		"routeros_bridge_vlan":         struct{}{},
		"routeros_gre":                 struct{}{},
		"routeros_ipip":                struct{}{},
		"routeros_vlan":                struct{}{},
		"routeros_vrrp":                struct{}{},
		"routeros_wireguard":           struct{}{},
		"routeros_wireguard_peer":      struct{}{},
		"routeros_identity":            struct{}{},
		"routeros_scheduler":           struct{}{},
	}
)
