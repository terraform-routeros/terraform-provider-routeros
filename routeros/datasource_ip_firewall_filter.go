package routeros

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func getIPFirewallFilterSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				KeyFilter: PropFilterRw,
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"action": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address_list_timeout": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"bytes": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"chain": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyComment: {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_bytes": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_nat_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_rate": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"content": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyDisabled: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"dscp": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"dst_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_address_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_address_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyDynamic: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"fragment": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"hotspot": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"icmp_options": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_bridge_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_bridge_port_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_interface": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_interface_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ingress_priority": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"invalid": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"ipsec_policy": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ipv4_options": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"jump_target": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"hw_offload": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"layer7_protocol": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"log": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"log_prefix": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"nth": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_bridge_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_bridge_port_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_interface": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_interface_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"packets": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"packet_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"packet_size": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"per_connection_classifier": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"priority": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"protocol": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"psd": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"random": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"reject_with": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"routing_table": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"routing_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_mac_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tcp_flags": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tcp_mss": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"time": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tls_host": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ttl": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}
