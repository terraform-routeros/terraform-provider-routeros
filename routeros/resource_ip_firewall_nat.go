package routeros

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
 {
    ".id": "*2",
    "action": "masquerade",
    "bytes": "0",
    "chain": "srcnat",
    "disabled": "false",
    "dynamic": "false",
    "invalid": "false",
    "log": "false",
    "log-prefix": "",
    "out-interface": "all-wireless",
    "packets": "0",
    "src-address-list": "LAN"
  }
*/

// ResourceIPFirewallNat https://help.mikrotik.com/docs/display/ROS/NAT
func ResourceIPFirewallNat() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/nat"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("bytes", "packets"),
		MetaSetUnsetFields: PropSetUnsetFields("dst_address_list", "src_address_list", "in_interface", "in_interface_list",
			"out_interface", "out_interface_list", "in_bridge_port_list", "out_bridge_port_list"),

		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Action to take if a packet is matched by the rule",
			ValidateFunc: validation.StringInSlice([]string{
				"accept", "add-dst-to-address-list", "add-src-to-address-list", "dst-nat", "endpoint-independent-nat",
				"jump", "log", "masquerade", "netmap", "passthrough", "redirect", "return", "same", "src-nat",
			}, false),
		},
		"address_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the address list to be used. Applicable if action is add-dst-to-address-list or " +
				"add-src-to-address-list.",
		},
		"address_list_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time interval after which the address will be removed from the address list specified by " +
				"address-list parameter. Used in conjunction with add-dst-to-address-list or add-src-to-address-list " +
				"actions.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Specifies to which chain rule will be added. If the input does not match the name of an " +
				"already defined chain, a new chain will be created.",
		},
		KeyComment: PropCommentRw,
		"connection_bytes": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets only if a given amount of bytes has been transfered through the particular " +
				"connection.",
		},
		"connection_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches connections per address or address block after given value is reached. Should be " +
				"used together with connection-state=new and/or with tcp-flags=syn because matcher is very resource " +
				"intensive.",
		},
		"connection_mark": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets marked via mangle facility with particular connection mark. If no-mark is " +
				"set, rule will match any unmarked connection.",
		},
		// See comment for the "path_cost" field in resource_interface_bridge_port.go file.
		"connection_rate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Connection Rate is a firewall matcher that allow to capture traffic based on present speed " +
				"of the connection (0..4294967295).",
		},
		"connection_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets from related connections based on information from their connection " +
				"tracking helpers.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{
				"ftp", "h323", "irc", "pptp", "quake3", "sip", "tftp",
			}, false, true),
		},
		"content": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Match packets that contain specified text.",
		},
		KeyDisabled: PropDisabledRw,
		"dscp": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches DSCP IP header field.",
			ValidateFunc: validation.IntBetween(0, 63),
		},
		"dst_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets which destination is equal to specified IP or falls into specified IP range.",
		},
		"dst_address_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches destination address of a packet against user-defined address list.",
		},
		"dst_address_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matches destination address type.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"unicast", "local", "broadcast", "multicast"}, false, true),
		},
		"dst_limit": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets until a given rate is exceeded.",
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of destination port numbers or port number ranges.",
		},
		KeyDynamic: PropDynamicRo,
		"fragment": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Matches fragmented packets. First (starting) fragment does not count. If connection tracking " +
				"is enabled there will be no fragments as system automatically assembles every packet",
		},
		"hotspot": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matches packets received from HotSpot clients against various HotSpot matchers.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"auth", "from-client", "http", "local-dst", "to-client"}, false, true),
		},
		"icmp_options": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches ICMP type: code fields.",
		},
		"in_bridge_port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Actual interface the packet has entered the router if the incoming interface is a bridge. " +
				"Works only if use-ip-firewall is enabled in bridge settings.",
		},
		"in_bridge_port_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as in-bridge-port.",
		},
		"in_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface the packet has entered the router.",
		},
		"in_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as in-interface.",
		},
		"ingress_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Matches the priority of an ingress packet. Priority may be derived from VLAN, WMM, DSCP, " +
				"or MPLS EXP bit.",
			ValidateFunc: validation.IntBetween(0, 63),
		},
		"invalid": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"ipsec_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches the policy used by IPsec. Value is written in the following format: direction, policy.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(in|out)\s?,\s?(ipsec|none)$`),
				"Value must be written in the following format: direction, policy."),
		},
		"ipv4_options": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches IPv4 header options.",
			ValidateFunc: validation.StringInSlice([]string{
				"any", "loose-source-routing", "no-record-route", "no-router-alert", "no-source-routing",
				"no-timestamp", "none", "record-route", "router-alert", "strict-source-routing", "timestamp",
			}, false),
		},
		"jump_target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the target chain to jump to. Applicable only if action=jump.",
		},
		"layer7_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Layer7 filter name.",
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
		"nth": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches every nth packet: nth=2,1 rule will match every first packet of 2, hence, 50% of " +
				"all the traffic that is matched by the rule",
		},
		"out_bridge_port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Actual interface the packet is leaving the router if the outgoing interface is a bridge. " +
				"Works only if use-ip-firewall is enabled in bridge settings.",
		},
		"out_bridge_port_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as out-bridge-port.",
		},
		"out_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface the packet is leaving the router.",
		},
		"out_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as out-interface.",
		},
		"packet_mark": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets marked via mangle facility with particular packet mark. If no-mark is set, " +
				"the rule will match any unmarked packet.",
		},
		"packet_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets of specified size or size range in bytes.",
		},
		"per_connection_classifier": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "PCC matcher allows dividing traffic into equal streams with the ability to keep packets " +
				"with a specific set of options in one particular stream.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches if any (source or destination) port matches the specified list of ports or port " +
				"ranges. Applicable only if protocol is TCP or UDP",
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Matches the packet's priority after a new priority has been set. Priority may be derived from " +
				"VLAN, WMM, DSCP, MPLS EXP bit, or from the priority that has been set using the set-priority action.",
			ValidateFunc: validation.IntBetween(0, 63),
		},
		"protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches particular IP protocol specified by protocol name or number.",
		},
		"psd": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Attempts to detect TCP and UDP scans. Parameters are in the following format WeightThreshold, " +
				"DelayThreshold, LowPortWeight, HighPortWeight.",
		},
		"random": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches packets randomly with a given probability.",
			ValidateFunc: validation.IntBetween(1, 99),
		},
		"randomise_ports": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Randomize to which public port connections will be mapped.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"routing_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets marked by mangle facility with particular routing mark.",
		},
		"same_not_by_dst": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether to take into account or not destination IP address when selecting a " +
				"new source IP address. Applicable if action=same",
		},
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets which source is equal to specified IP or falls into a specified IP range.",
		},
		"src_address_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches source address of a packet against user-defined address list.",
		},
		"src_address_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matches source address type.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"unicast", "local", "broadcast", "multicast"}, false, true),
		},
		"src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of source ports and ranges of source ports. Applicable only if a protocol is TCP or UDP.",
		},
		"src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Matches source MAC address of the packet.",
			ValidateFunc: validation.IsMACAddress,
		},
		"tcp_mss": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches TCP MSS value of an IP packet.",
		},
		"time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows to create a filter based on the packets' arrival time and date or, for locally " +
				"generated packets, departure time and date.",
		},
		"to_addresses": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Replace original address with specified one. Applicable if action is dst-nat, netmap, " +
				"same, src-nat.",
		},
		"to_ports": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Replace the original port with the specified one. Applicable if action is dst-nat, " +
				"redirect, masquerade, netmap, same, src-nat.",
		},
		"ttl": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches packets TTL value.",
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
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
