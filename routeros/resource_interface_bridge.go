package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceBridge https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge
func ResourceInterfaceBridge() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge"),
		MetaId:           PropId(Id),

		KeyActualMtu: PropActualMtuRo,
		"add_dhcp_option82": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to add DHCP Option-82 information (Agent Remote ID and Agent Circuit ID) to DHCP " +
				"packets. Can be used together with Option-82 capable DHCP server to assign IP addresses and implement " +
				"policies. This property only has effect when dhcp-snooping is set to yes.",
			RequiredWith: []string{"dhcp_snooping"},
		},
		"admin_mac": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "Static MAC address of the bridge. This property only has effect when auto-mac is set to no.",
		},
		"ageing_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How long a host's information will be kept in the bridge database.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"auto_mac": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Automatically select one MAC address of bridge ports as a bridge MAC address, bridge MAC " +
				"will be chosen from the first added bridge port. After a device reboot, the bridge MAC " +
				"can change depending on the port-number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"dhcp_snooping": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		KeyDisabled: PropDisabledRw,
		"ether_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "This property only has effect when vlan-filtering is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"0x9100", "0x8100", "0x88a8"}, false),
		},
		"fast_forward": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"forward_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time which is spent during the initialization phase of the bridge interface " +
				"(i.e., after router startup or enabling the interface) in listening/learning state before the " +
				"bridge will start functioning normally.",
			DiffSuppressFunc: TimeEquall,
		},
		"frame_types": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies allowed frame types on a bridge port. This property only has effect when " +
				"vlan-filtering is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc: validation.StringInSlice([]string{"admit-all", "admit-only-untagged-and-priority-tagged",
				"admit-only-vlan-tagged"}, false),
		},
		"igmp_snooping": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables multicast group and port learning to prevent multicast traffic from flooding all " +
				"interfaces in a bridge.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"igmp_version": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "Selects the IGMP version in which IGMP general membership queries will be generated. " +
				"This property only has effect when igmp-snooping is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntInSlice([]int{2, 3}),
			RequiredWith:     []string{"igmp_snooping"},
		},
		"ingress_filtering": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables or disables VLAN ingress filtering, which checks if the ingress port is a member " +
				"of the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if " +
				"the ingress traffic should be tagged or untagged. This property only has effect when " +
				"vlan-filtering is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			RequiredWith:     []string{"vlan_filtering"},
		},
		KeyL2Mtu: PropL2MtuRo,
		"last_member_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "If a port has fast-leave set to no and a bridge port receives a IGMP Leave message, " +
				"then a IGMP Snooping enabled bridge will send a IGMP query to make sure that no devices has " +
				"subscribed to a certain multicast stream on a bridge port.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping"},
		},
		"last_member_query_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "How many times should last-member-interval pass until a IGMP Snooping bridge will stop " +
				"forwarding a certain multicast stream. This property only has effect when igmp-snooping is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			RequiredWith:     []string{"igmp_snooping"},
		},
		KeyMacAddress: PropMacAddressRo,
		"max_hops": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "Bridge count which BPDU can pass in a MSTP enabled network in the same region before BPDU " +
				"is being ignored. This property only has effect when protocol-mode is set to mstp.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(6, 40),
		},
		"max_message_age": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Changes the Max Age value in BPDU packets, which is transmitted by the root bridge. " +
				"This property only has effect when protocol-mode is set to stp or rstp. Value: 6s..40s",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"membership_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Amount of time after an entry in the Multicast Database (MDB) is removed if a IGMP membership " +
				"report is not received on a certain port. This property only has effect when igmp-snooping is set to yes.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping"},
		},
		"mld_version": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "Selects the MLD version. Version 2 adds support for source-specific multicast. This " +
				"property only has effect when RouterOS IPv6 package is enabled and igmp-snooping is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntInSlice([]int{1, 2}),
			RequiredWith:     []string{"igmp_snooping"},
		},
		"mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The default bridge MTU value without any bridge ports added is 1500. " +
				"The MTU value can be set manually, but it cannot exceed the bridge L2MTU or the lowest bridge " +
				"port L2MTU. If a new bridge port is added with L2MTU which is smaller than the actual-mtu " +
				"of the bridge (set by the mtu property), then manually set value will be ignored and the bridge " +
				"will act as if mtu=auto is set.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multicast_querier": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Multicast querier generates IGMP general membership queries to which all IGMP capable " +
				"devices respond with an IGMP membership report, usually a PIM (multicast) router or IGMP proxy " +
				"generates these queries. This property only has an effect when igmp-snooping is set to yes. " +
				"Additionally, the igmp-snooping should be disabled/enabled after changing multicast-querier property.",
			RequiredWith: []string{"igmp_snooping"},
		},
		//  https://help.mikrotik.com/docs/pages/viewpage.action?pageId=59277403#BridgeIGMP/MLDsnooping-Configurationoptions
		"multicast_router": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "A multicast router port is a port where a multicast router or querier is connected. On " +
				"this port, unregistered multicast streams and IGMP/MLD membership reports will be sent. This " +
				"setting changes the state of the multicast router for a bridge interface itself. This property can " +
				"be used to send IGMP/MLD membership reports and multicast traffic to the bridge interface for further " +
				"multicast routing or proxying. This property only has an effect when igmp-snooping is set to yes.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "permanent", "temporary-query"}, false),
			RequiredWith: []string{"igmp_snooping"},
		},
		"mvrp": {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:          false,
			Description:      "Enables MVRP for bridge (available since RouterOS 7.15). It ensures that the MAC address 01:80:C2:00:00:21 is trapped and not forwarded, the vlan-filtering must be enabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropNameForceNewRw,
		"port_cost_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option that changes the port path cost and internal path cost mode for bridged ports, utilizing automatic values based on interface speed.",
			ValidateFunc:     validation.StringInSlice([]string{"long", "short"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Bridge priority, used by STP to determine root bridge, used by MSTP to determine CIST " +
				"and IST regional root bridge. This property has no effect when protocol-mode is set to none.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"protocol_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Select Spanning tree protocol (STP) or Rapid spanning tree protocol (RSTP) to ensure a " +
				"loop-free topology for any bridged LAN.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"none", "rstp", "stp", "mstp"}, false),
		},
		"pvid": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Port VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. " +
				"It applies e.g. to frames sent from bridge IP and destined to a bridge port. " +
				"This property only has effect when vlan-filtering is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(1, 4094),
		},
		"querier_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Used to change the interval how often a bridge checks if it is the active multicast " +
				"querier. This property only has effect when igmp-snooping and multicast-querier is set to yes.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping", "multicast_querier"},
		},
		"query_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Used to change the interval how often IGMP general membership queries are sent out. " +
				"This property only has effect when igmp-snooping and multicast-querier is set to yes.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping", "multicast_querier"},
		},
		"query_response_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Interval in which a IGMP capable device must reply to a IGMP query with a IGMP membership " +
				"report. This property only has effect when igmp-snooping and multicast-querier is set to yes.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping", "multicast_querier"},
		},
		KeyRunning: PropRunningRo,
		"region_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "MSTP region name. This property only has effect when protocol-mode is set to mstp.",
		},
		"region_revision": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "MSTP configuration revision number. This property only has effect when protocol-mode is set to mstp.",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"startup_query_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "Specifies how many times must startup-query-interval pass until the bridge starts sending " +
				"out IGMP general membership queries periodically. This property only has effect when igmp-snooping " +
				"and multicast-querier is set to yes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			RequiredWith:     []string{"igmp_snooping", "multicast_querier"},
		},
		"startup_query_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Used to change the amount of time after a bridge starts sending out IGMP general membership " +
				"queries after the bridge is enabled. This property only has effect when igmp-snooping and " +
				"multicast-querier is set to yes.",
			DiffSuppressFunc: TimeEquall,
			RequiredWith:     []string{"igmp_snooping", "multicast_querier"},
		},
		"transmit_hold_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The Transmit Hold Count used by the Port Transmit state machine to limit transmission rate.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(1, 10),
		},
		"vlan_filtering": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Globally enables or disables VLAN functionality for bridge.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// Some properties are not implemented, see: https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceInterfaceBridgeV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
