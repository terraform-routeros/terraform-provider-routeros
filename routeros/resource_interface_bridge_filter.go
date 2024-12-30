package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*3C",
  "action": "drop",
  "arp-dst-mac-address": "00:00:00:00:00:00/FF:FF:FF:FF:FF:FF",
  "arp-gratuitous": "false",
  "arp-hardware-type": "0",
  "arp-opcode": "0",
  "arp-packet-type": "0",
  "arp-src-mac-address": "00:00:00:00:00:00/FF:FF:FF:FF:FF:FF",
  "bytes": "0",
  "chain": "vpn_to_hsia_bridge",
  "comment": "Explicit block",
  "disabled": "false",
  "dst-mac-address": "00:00:00:00:00:00/FF:FF:FF:FF:FF:FF",
  "dynamic": "false",
  "in-bridge": "bridge_cast",
  "in-bridge-list": "dynamic",
  "ingress-priority": "0",
  "invalid": "false",
  "limit": "1,5",
  "log": "true",
  "log-prefix": "HSIA -> VPN Blocked ",
  "mac-protocol": "arp",
  "out-bridge": "bridge_cast",
  "out-bridge-list": "static",
  "packet-mark": "mdns_to_be_dropped",
  "packet-type": "broadcast",
  "packets": "0",
  "src-mac-address": "00:00:00:00:00:00/FF:FF:FF:FF:FF:FF"
}

*/

// ResourceIPFirewallFilter https://help.mikrotik.com/docs/spaces/ROS/pages/328068/Bridging+and+Switching#BridgingandSwitching-BridgeFirewall
func ResourceInterfaceBridgeFilter() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/filter"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("bytes", "packets", "invalid"),
		MetaSetUnsetFields: PropSetUnsetFields("arp_dst_mac_address", "arp_gratuitous", "arp_hardware_type",
			"arp_opcode", "arp_packet_type", "arp_src_address", "arp_src_mac_address", "dst_address", "dst_mac_address",
			"dst_port", "in_bridge", "in_bridge_list", "in_interface", "in_interface_list", "ingress_priority",
			"ip_protocol", "limit", "mac_protocol", "new_packet_mark", "new_priority", "out_bridge", "out_bridge_list",
			"out_interface", "out_interface_list", "packet_mark", "packet_type", "src_address", "src_mac_address",
			"src_port", "stp_flags", "stp_forward_delay", "stp_hello_time", "stp_max_age", "stp_root_address",
			"stp_port", "stp_root_cost", "stp_root_priority", "stp_sender_address", "stp_sender_priority", "stp_type",
			"tls_host", "vlan_encap", "vlan_id", "vlan_priority"),

		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Action to take if a packet is matched by the rule",
			ValidateFunc: validation.StringInSlice([]string{
				"accept", "drop", "mark-packet", "jump", "log", "passthrough", "set-priority", "return",
			}, false),
		},
		"arp_dst_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP destination MAC address",
			ValidateFunc: ValidationMacAddress,
		},
		"arp_gratuitous": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Matches ARP gratuitous packets.",
		},
		"arp_hardware_type": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "ARP hardware type. This is normally Ethernet (Type 1).",
		},
		"arp_opcode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Action to take if a packet is matched by the rule",
			ValidateFunc: validation.StringInSlice([]string{
				"arp-nak", "drarp-error", "rarp-reply", "drarp-request", "inarp-reply",
				"inarp-request", "reply", "reply-reverse", "request", "request-reverse",
			}, true),
		},
		"arp_packet_type": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "ARP Packet Type",
			ValidateFunc: Validation64k,
		},
		"arp_src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP source IP address.",
			ValidateFunc: ValidationIpAddress,
		},
		"arp_src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP source MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Specifies to which chain rule will be added. If the input does not match the name of an " +
				"already defined chain, a new chain will be created.",
		},
		KeyComment: PropCommentRw,
		"dst_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Destination IP address (only if MAC protocol is set to IP).",
			ValidateFunc: ValidationIpAddress,
		},
		"dst_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Destination MAC address.",
			ValidateFunc: ValidationMacAddressWithMask,
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of destination port numbers or port number ranges.",
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"in_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge interface through which the packet is coming in.",
		},
		"in_bridge_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of bridge interfaces defined in interface list. Works the same as in-bridge.",
		},
		"in_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Physical interface (i.e., bridge port) through which the packet is coming in.",
		},
		"in_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as in-interface.",
		},
		"ingress_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Integer. Matches the priority of an ingress packet. Priority may be derived from VLAN, WMM, DSCP,or MPLS EXP bit.",
			ValidateFunc: Validation64k,
		},
		"ip_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP protocol (only if MAC protocol is set to IPv4)",
			ValidateFunc: validation.StringInSlice([]string{
				"dccp", "egp", "encap", "etherip", "ggp", "gre", "hmp", "icmp", "icmpv6",
				"idpr-cmtp", "igmp", "ipencap", "ipip", "ipsec-ah", "ipsec-esp", "ipv6",
				"ipv6-frag", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "iso-tp4", "l2tp",
				"ospf", "pim", "pup", "rdp", "rspf", "rsvp", "sctp", "st", "tcp", "udp", "udp-lite",
				"vmtp", "vrrp", "xns-idp", "xtp",
			}, true),
		},
		"jump_target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the target chain to jump to. Applicable only if action=jump.",
		},
		"limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets up to a limited rate (packet rate or bit rate). A rule using this matcher " +
				"will match until this limit is reached. Parameters are written in the following format: " +
				"rate[/time],burst:mode.",
		},
		"log": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Add a message to the system log.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"log_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adds specified text at the beginning of every log message. Applicable if action=log or " +
				"log=yes configured.",
		},
		"mac_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Ethernet payload type (MAC-level protocol). To match protocol type for VLAN encapsulated " +
				"frames (0x8100 or 0x88a8), a vlan-encap property should be used.",
			ValidateFunc: validation.StringInSlice([]string{
				"802.2", "arp", "homeplug-av", "ip", "ipv6", "ipx", "length",
				"lldp", "loop-protect", "mpls-multicast", "mpls-unicast", "packing-compr",
				"packing-simple", "pppoe", "pppoe-discovery", "rarp", "service-vlan", "vlan",
			}, true),
		},
		"new_packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets a new packet-mark value.",
		},
		"new_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Sets a new priority for a packet. This can be the VLAN, WMM or MPLS EXP priority",
			ValidateFunc: Validation64k,
		},
		"out_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge interface through which the packet going out.",
		},
		"out_bridge_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of bridge interfaces defined in interface list. Works the same as out-bridge.",
		},
		"out_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface the packet has entered the router.",
		},
		"out_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as out-interface.",
		},
		"packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Match packets with a certain packet mark.",
		},
		"packet_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Match packets with a certain packet mark.",
			ValidateFunc: validation.StringInSlice([]string{"broadcast", "host", "multicast", "other-host"}, true),
		},
		"passthrough": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to let the packet to pass further (like action passthrough) into the " +
				"filter or not (property only valid some actions).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyPlaceBefore: PropPlaceBefore,
		"src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source port number or range (only for TCP or UDP protocols).",
			ValidateFunc: ValidationIpAddress,
		},
		"src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of source port numbers or port number ranges.",
		},
		"stp_flags": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Match packets with a certain packet mark.",
			ValidateFunc: validation.StringInSlice([]string{"topology-change", "topology-change-ack"}, false),
		},
		"stp_forward_delay": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Forward delay timer.",
			ValidateFunc: Validation64k,
		},
		"stp_hello_time": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP hello packets time.",
			ValidateFunc: Validation64k,
		},
		"stp_max_age": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Maximal STP message age.",
			ValidateFunc: Validation64k,
		},
		"stp_root_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Root bridge MAC address",
			ValidateFunc: ValidationMacAddress,
		},
		"stp_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP port identifier.",
			ValidateFunc: Validation64k,
		},
		"stp_root_cost": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Root bridge cost.",
			ValidateFunc: Validation64k,
		},

		"stp_root_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "",
			ValidateFunc: Validation64k,
		},
		"stp_sender_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "STP message sender MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"stp_sender_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP sender priority.",
			ValidateFunc: Validation64k,
		},
		"stp_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The BPDU type: config - configuration BPDU OR tcn - topology change notification",
			ValidateFunc: validation.StringInSlice([]string{"config", "tcn"}, false),
		},
		"tls_host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Allows matching https traffic based on TLS SNI hostname. Accepts GLOB syntax for wildcard matching",
		},
		"vlan_encap": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches the MAC protocol type encapsulated in the VLAN frame.",
			ValidateFunc: Validation64k,
		},
		"vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches the VLAN identifier field.",
			ValidateFunc: validation.IntBetween(0, 4095),
		},
		"vlan_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches the VLAN identifier field.",
			ValidateFunc: validation.IntBetween(0, 7),
		},
	}
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			skip := resSchema[MetaSkipFields].Default.(string)
			resSchema[MetaSkipFields].Default = skip + `,"place_before"`
			defer func() {
				resSchema[MetaSkipFields].Default = skip
			}()

			return ResourceUpdate(ctx, resSchema, d, m)
		},
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
