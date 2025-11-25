package main

var (
	providerAliasesMap = map[string]struct{}{
		"routeros_dhcp_client":         {},
		"routeros_dhcp_client_option":  {},
		"routeros_dhcp_server":         {},
		"routeros_dhcp_server_network": {},
		"routeros_dhcp_server_lease":   {},
		"routeros_firewall_addr_list":  {},
		"routeros_firewall_filter":     {},
		"routeros_firewall_mangle":     {},
		"routeros_firewall_nat":        {},
		"routeros_dns":                 {},
		"routeros_dns_record":          {},
		"routeros_bridge":              {},
		"routeros_bridge_mlag":         {},
		"routeros_bridge_port":         {},
		"routeros_bridge_vlan":         {},
		"routeros_gre":                 {},
		"routeros_ipip":                {},
		"routeros_vlan":                {},
		"routeros_vrrp":                {},
		"routeros_wireguard":           {},
		"routeros_wireguard_peer":      {},
		"routeros_identity":            {},
		"routeros_scheduler":           {},
	}
)
