package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
     {
    ".id": "*1",
    "add-path-out": "none",
    "address-families": "ip",
    "as": "65521/1",
    "as-override": "true",
    "cisco-vpls-nlri-len-fmt": "auto-bits",
    "cluster-id": "0.0.0.0",
    "connect": "true",
    "disabled": "false",
    "hold-time": "infinity",
    "inactive": "false",
    "input.accept-communities": "",
    "input.accept-ext-communities": "",
    "input.accept-large-communities": "",
    "input.accept-nlri": "",
    "input.accept-unknown": "",
    "input.affinity": "alone",
    "input.allow-as": "0",
    "input.filter": "",
    "input.ignore-as-path-len": "true",
    "input.limit-process-routes-ipv4": "5",
    "input.limit-process-routes-ipv6": "5",
    "keepalive-time": "3m",
    "listen": "true",
    "local.address": "127.0.0.1",
    "local.port": "22334",
    "local.role": "ibgp",
    "local.ttl": "3",
    "multihop": "true",
    "name": "bgp1",
    "nexthop-choice": "default",
    "output.affinity": "alone",
    "output.default-originate": "never",
    "output.default-prepend": "1",
    "output.filter-chain": "",
    "output.filter-select": "",
    "output.keep-sent-attributes": "true",
    "output.network": "",
    "output.no-client-to-client-reflection": "false",
    "output.no-early-cut": "true",
    "output.redistribute": "rip,bgp",
    "remote.address": "0.0.0.0/32",
    "remote.allowed-as": "1111",
    "remote.port": "11223",
    "remote.ttl": "3",
    "remove-private-as": "true",
    "router-id": "0.0.0.1",
    "routing-table": "main",
    "save-to": "bgp.dump",
    "tcp-md5-key": "poipoipoipoipoi",
    "templates": "test-template",
    "use-bfd": "true",
    "vrf": "main"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/BGP#BGP-ConnectionMenu
func ResourceRoutingBGPConnection() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bgp/connection"),
		MetaId:           PropId(Id),

		"add_path_out": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			Default:      "none",
			ValidateFunc: validation.StringInSlice([]string{"all", "none"}, false),
		},
		"address_families": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ip",
			Description: "List of address families about which this peer will exchange routing information. The " +
				"remote peer must support (they usually do) BGP capabilities optional parameter to " +
				"negotiate any other families than IP.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"ip", "ipv6", "l2vpn", "l2vpn-cisco", "vpnv4"}, false, false),
		},
		"as": {
			Type:     schema.TypeString,
			Required: true,
			Description: "32-bit BGP autonomous system number. Value can be entered in AS-Plain and AS-Dot " +
				"formats. The parameter is also used to set up the BGP confederation, in the following " +
				"format: confederation_as/as . For example, if your AS is 34 and your confederation AS is " +
				"43, then as configuration should be as =43/34.",
		},
		"cisco_vpls_nlri_len_fmt": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "VPLS NLRI length format type. Used for compatibility with Cisco VPLS.",
			ValidateFunc: validation.StringInSlice([]string{"auto-bits", "auto-bytes", "bits", "bytes"}, false),
		},
		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "In case this instance is a route reflector: the cluster ID of the router reflector " +
				"cluster to this instance belongs. This attribute helps to recognize routing updates " +
				"that come from another route reflector in this cluster and avoid routing information " +
				"looping. Note that normally there is only one route reflector in a cluster; in this " +
				"case, 'cluster-id' does not need to be configured and BGP router ID is used instead.",
			ValidateFunc: validation.IsIPv4Address,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"connect": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to allow the router to initiate the connection.",
		},
		"hold_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Specifies the BGP Hold Time value to use when negotiating with peers. According to the " +
				"BGP specification, if the router does not receive successive KEEPALIVE and/or UPDATE " +
				"and/or NOTIFICATION messages within the period specified in the Hold Time field of the " +
				"OPEN message, then the BGP connection to the peer will be closed. The minimal hold-time " +
				"value of both peers will be actually used (note that the special value 0 or 'infinity' " +
				"is lower than any other value) infinity - never expire the connection and never send " +
				"keepalive messages.",
		},
		"inactive": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"input": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with BGP input.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"accept_communities": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "A quick way to filter incoming updates with specific communities. It allows filtering " +
							"incoming messages directly before they are even parsed and stored in memory, that way " +
							"significantly reducing memory usage. Regular input filter chain can only reject " +
							"prefixes which means that it will still eat memory and will be visible in /routing " +
							"route table as 'not active, filtered'. Changes to be applied required session refresh.",
					},
					"accept_ext_communities": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "A quick way to filter incoming updates with specific extended communities. It allows " +
							"filtering incoming messages directly before they are even parsed and stored in memory, " +
							"that way significantly reducing memory usage. Regular input filter chain can only " +
							"reject prefixes which means that it will still eat memory and will be visible in " +
							"/routing route table as 'not active, filtered'. Changes to be applied required session " +
							"refresh.",
					},
					"accept_large_communities": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "A quick way to filter incoming updates with specific large communities. It allows " +
							"filtering incoming messages directly before they are even parsed and stored in memory, " +
							"that way significantly reducing memory usage. Regular input filter chain can only " +
							"reject prefixes which means that it will still eat memory and will be visible in " +
							"/routing route table as 'not active, filtered'. Changes to be applied required session " +
							"refresh.",
					},
					"accept_nlri": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Name of the ipv4/6 address-list. A quick way to filter incoming updates with specific " +
							"NLRIs. It allows filtering incoming messages directly before they are even parsed and " +
							"stored in memory, that way significantly reducing memory usage. Regular input filter " +
							"chain can only reject prefixes which means that it will still eat memory and will be " +
							"visible in /routing route table as 'not active, filtered'. Changes to be applied " +
							"required session restart.",
					},
					"accept_unknown": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "A quick way to filter incoming updates with specific 'unknown' attributes. It allows " +
							"filtering incoming messages directly before they are even parsed and stored in memory, " +
							"that way significantly reducing memory usage. Regular input filter chain can only " +
							"reject prefixes which means that it will still eat memory and will be visible in " +
							"/routing route table as 'not active, filtered'. Changes to be applied required session " +
							"refresh.",
					},
					// affinity (afi | alone | instance | main | remote-as | vrf; Default: )
					// May be "0"
					"affinity": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Configure input multi-core processing. Read more in Routing Protocol Multi-core Support " +
							"article. alone - input and output of each session are processed in its own process, " +
							"most likely the best option when there are a lot of cores and a lot of peers afi, " +
							"instance, vrf, remote-as - try to run input/output of new session in process with " +
							"similar parameters main - run input/output in the main process (could potentially " +
							"increase performance on single-core even possibly on multi-core devices with a small " +
							"amount of cores) input - run output in the same process as input (can be set only for " +
							"output affinity)",
					},
					"allow_as": {
						Type:     schema.TypeInt,
						Optional: true,
						Description: "Indicates how many times to allow your own AS number in AS-PATH, before discarding a " +
							"prefix.",
						ValidateFunc: validation.IntBetween(0, 10),
					},
					"filter": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Name of the routing filter chain to be used on input prefixes. This happens after " +
							"NLRIs are processed. If the chain is not specified, then BGP by default accepts " +
							"everything.",
					},
					"ignore_as_path_len": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether to ignore the AS_PATH attribute in the BGP route selection algorithm",
					},
					// FIXME ROS 7.8: 'unknown parameter input.limit-nlri-diversity'
					// "limit_nlri_diversity": {
					// 	Type:     schema.TypeInt,
					// 	Optional: true,
					// },
					"limit_process_routes_ipv4": {
						Type:     schema.TypeInt,
						Optional: true,
						Description: "Try to limit the amount of received IPv4 routes to the specified number. This number " +
							"does not represent the exact number of routes going to be installed in the routing " +
							"table by the peer. BGP session 'clear' command must be used to reset the flag if the " +
							"limit is reached.",
					},
					"limit_process_routes_ipv6": {
						Type:     schema.TypeInt,
						Optional: true,
						Description: "Try to limit the amount of received IPv6 routes to the specified number. This number " +
							"does not represent the exact number of routes going to be installed in the routing " +
							"table by the peer. BGP session 'clear' command must be used to reset the flag if the " +
							"limit is reached.",
					},
				},
			},
		},
		"keepalive_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "3m",
			Description:      "How long to keep the BGP session open after the last received 'keepalive' message.",
			DiffSuppressFunc: TimeEquall,
		},
		"listen": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to listen for incoming connections.",
		},
		"local": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with BGP input.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"address": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Local connection IPv4/6 address.",
					},
					"default_address": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "",
					},
					"port": {
						Type:         schema.TypeInt,
						Optional:     true,
						Default:      179,
						Description:  "Local connection port.",
						ValidateFunc: Validation64k,
					},
					"role": {
						Type:     schema.TypeString,
						Required: true,
						Description: "BGP role, in most common scenarios it should be set to iBGP or eBGP. More " +
							"information on BGP roles can be found in the corresponding [RFC draft]" +
							"(https://datatracker.ietf.org/doc/draft-ietf-idr-bgp-open-policy/?include_text=1)",
						ValidateFunc: validation.StringInSlice(
							[]string{
								"ebgp",
								"ebgp-customer",
								"ebgp-peer",
								"ebgp-provider",
								"ebgp-rs",
								"ebgp-rs-client",
								"ibgp",
								"ibgp-rr",
								"ibgp-rr-client",
							},
							false,
						),
					},
					"ttl": {
						Type:         schema.TypeInt,
						Optional:     true,
						Description:  "Time To Live (hop limit) that will be recorded in sent TCP packets.",
						ValidateFunc: validation.IntBetween(1, 255),
					},
				},
			},
		},
		"multihop": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether the remote peer is more than one hop away. This option affects " +
				"outgoing next-hop selection as described in RFC 4271 (for EBGP only, excluding EBGP " +
				"peers local to the confederation). It also affects: whether to accept connections from " +
				"peers that are not in the same network (the remote address of the connection is used " +
				"for this check); whether to accept incoming routes with NEXT_HOP attribute that is not " +
				"in the same network as the address used to establish the connection; the target-scope " +
				"of the routes installed from this peer; routes from multi-hop or IBGP peers resolve " +
				"their next-hops through IGP routes by default.",
		},
		KeyName: PropName("Name of the BGP connection."),
		"nexthop_choice": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Affects the outgoing NEXT_HOP attribute selection. Note that next-hops set in filters " +
				"always take precedence. Also note that the next-hop is not changed on route reflection, " +
				"except when it's set in the filter. default - select the next-hop as described in RFC " +
				"4271 force-self - always use a local address of the interface that is used to connect to " +
				"the peer as the next-hop; propagate - try to propagate further the next-hop received; " +
				"i.e. if the route has BGP NEXT_HOP attribute, then use it as the next-hop, otherwise, " +
				"fall back to the default case.",
			ValidateFunc: validation.StringInSlice([]string{"default", "force-self", "propagate"}, false),
		},
		"output": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with BGP output.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					// May be "0" ?!?
					// affinity (afi | alone | instance | main | remote-as | vrf; Default: )
					"affinity": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Configure output multicore processing. Read more in Routing Protocol Multi-core Support " +
							"article. alone - input and output of each session is processed in its own process, the " +
							"most likely best option when there are a lot of cores and a lot of peers afi, instance, " +
							"vrf, remote-as - try to run input/output of new session in process with similar " +
							"parameters main - run input/output in the main process (could potentially increase " +
							"performance on single-core even possibly on multicore devices with small amount of " +
							"cores) input - run output in the same process as input (can be set only for output " +
							"affinity).",
					},
					"as_override": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "If set, then all instances of the remote peer's AS number in the BGP AS-PATH attribute " +
							"are replaced with the local AS number before sending a route update to that peer. " +
							"Happens before routing filters and prepending.",
						DiffSuppressFunc: AlwaysPresentNotUserProvided,
					},
					"default_originate": {
						Type:         schema.TypeString,
						Optional:     true,
						Description:  "Specifies default route (0.0.0.0/0) distribution method.",
						ValidateFunc: validation.StringInSlice([]string{"always", "if-installed", "never"}, false),
					},
					"default_prepend": {
						Type:         schema.TypeInt,
						Optional:     true,
						Description:  "The count of AS prepended to the AS path.",
						ValidateFunc: validation.IntBetween(0, 255),
					},
					"filter_chain": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Name of the routing filter chain to be used on the output prefixes. If the chain is " +
							"not specified, then BGP by default accepts everything.",
					},
					"filter_select": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Name of the routing select chain to be used for prefix selection. If not specified, then " +
							"default selection is used.",
					},
					"keep_sent_attributes": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "Store in memory sent prefix attributes, required for ' dump-saved-advertisements ' " +
							"command to work. By default, sent-out prefixes are not stored to preserve the router's " +
							"memory. An option should be enabled only for debugging purposes when necessary to see " +
							"currently advertised prefixes.",
					},
					"network": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Name of the address list used to send local networks. The network is sent only if a " +
							"matching IGP route exists in the routing table.",
					},
					"no_client_to_client_reflection": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Disable client-to-client route reflection in Route Reflector setups.",
					},
					"no_early_cut": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "The early cut is the mechanism, to guess (based on default RFC behavior) what would " +
							"happen with the sent NPLRI when received by the remote peer. If the algorithm " +
							"determines that the NLRI is going to be dropped, a peer will not even try to send it. " +
							"However such behavior may not be desired in specific scenarios, then then this option " +
							"should be used to disable the early cut feature.",
					},
					"redistribute": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable redistribution of specified route types.",
						ValidateDiagFunc: ValidationMultiValInSlice([]string{
							"bgp", "connected", "bgp-mpls-vpn", "dhcp", "fantasy", "modem", "ospf", "rip", "static", "vpn",
						}, false, false),
					},
					"remove_private_as": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "If set, then the BGP AS-PATH attribute is removed before sending out route updates if " +
							"the attribute contains only private AS numbers. The removal process happens before " +
							"routing filters are applied and before the local, AS number is prepended to the AS path.",
						DiffSuppressFunc: AlwaysPresentNotUserProvided,
					},
				},
			},
		},
		"remote": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with BGP input.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"address": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Remote IPv4/6 address used to connect and/or listen to.",
					},
					"allowed_as": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "List of remote AS numbers that are allowed to connect. Useful for dynamic peer " +
							"configuration.",
					},
					"as": {
						Type:     schema.TypeString,
						Optional: true,
						Description: "Remote AS number. If not specified BGP will determine remote AS automatically " +
							"from the OPEN message.",
					},
					"port": {
						Type:         schema.TypeInt,
						Optional:     true,
						Description:  "Local connection port.",
						Default:      179,
						ValidateFunc: Validation64k,
					},
					"ttl": {
						Type:     schema.TypeInt,
						Optional: true,
						Description: "Acceptable minimum Time To Live, the hop limit for this TCP connection. For " +
							"example, if 'ttl=255' then only single-hop neighbors will be able to establish the " +
							"connection. This property only affects EBGP peers.",
						ValidateFunc: validation.IntBetween(1, 255),
					},
				},
			},
		},
		"router_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "BGP Router ID to be used. Use the ID from the /routing/router-id configuration by " +
				"specifying the reference name, or set the ID directly by specifying IP. Equal " +
				"router-ids are also used to group peers into one instance.",
		},
		"routing_table": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the routing table, to install routes in.",
		},
		"save_to": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Filename to be used to save BGP protocol-specific packet content (Exported PDU) into " +
				"pcap file. This method allows much simpler peer-specific packet capturing for debugging " +
				"purposes. Pcap files in this format can also be loaded to create virtual BGP peers to " +
				"recreate conditions that happened at the time when packet capture was running.",
		},
		"tcp_md5_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "The key used to authenticate the connection with TCP MD5 signature as described in RFC 2385. " +
				"If not specified, authentication is not used.",
		},
		"templates": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "List of the template names, to inherit parameters from. Useful for dynamic BGP peers.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"use_bfd": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to use the BFD protocol for faster connection state detection.",
		},
		KeyVrf: PropVrfRw,
	}

	return &schema.Resource{
		Description: "> [!WARNING] Using this resource you may happen unexpected behavior, for example, some of the attributes " +
			"may not be removable after adding them to the TF configuration. Please report this to GitHub and it " +
			"may be possible to fix it. Use the resource at your own risk as it is!",
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
