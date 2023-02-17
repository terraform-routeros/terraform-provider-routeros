package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceBridge https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge
func ResourceInterfaceBridge() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge"),
		MetaId:           PropId(Name),

		KeyActualMtu: PropActualMtuRo,
		"admin_mac": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "Static MAC address of the bridge. This property only has effect when auto-mac is set to no.",
		},
		"ageing_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "5m",
			Description: "How long a host's information will be kept in the bridge database.",
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"auto_mac": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "Automatically select one MAC address of bridge ports as a bridge MAC address, bridge MAC " +
				"will be chosen from the first added bridge port. After a device reboot, the bridge MAC " +
				"can change depending on the port-number.",
		},
		KeyComment: PropCommentRw,
		"dhcp_snooping": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		KeyDisabled: PropDisabledRw,
		"ether_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0x8100",
			Description:  "This property only has effect when vlan-filtering is set to yes.",
			ValidateFunc: validation.StringInSlice([]string{"0x9100", "0x8100", "0x88a8"}, false),
		},
		"fast_forward": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"forward_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "15s",
			Description: "Time which is spent during the initialization phase of the bridge interface " +
				"(i.e., after router startup or enabling the interface) in listening/learning state before the " +
				"bridge will start functioning normally.",
		},
		"frame_types": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "admit-all",
			Description: "Specifies allowed frame types on a bridge port. This property only has effect when " +
				"vlan-filtering is set to yes.",
			ValidateFunc: validation.StringInSlice([]string{"admit-all", "admit-only-untagged-and-priority-tagged",
				"admit-only-vlan-tagged"}, false),
		},
		"igmp_snooping": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Enables multicast group and port learning to prevent multicast traffic from flooding all " +
				"interfaces in a bridge.",
		},
		"igmp_version": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "If IGMP snooping is enabled, this sets the version of IGMP to use",
		},
		"ingress_filtering": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Enables or disables VLAN ingress filtering, which checks if the ingress port is a member " +
				"of the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if " +
				"the ingress traffic should be tagged or untagged. This property only has effect when " +
				"vlan-filtering is set to yes.",
		},
		KeyL2Mtu: PropL2MtuRo,
		"last_member_interval": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If IGMP snooping is enabled, this is the interval that the last member was seen.",
		},
		"last_member_query_count": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If IGMP snooping is enabled, this is the query count of the last member.",
		},
		"mac_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"max_message_age": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "20s",
			Description: "Changes the Max Age value in BPDU packets, which is transmitted by the root bridge. " +
				"This property only has effect when protocol-mode is set to stp or rstp. Value: 6s..40s",
		},
		"membership_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "auto",
			Description: "The default bridge MTU value without any bridge ports added is 1500. " +
				"The MTU value can be set manually, but it cannot exceed the bridge L2MTU or the lowest bridge " +
				"port L2MTU. If a new bridge port is added with L2MTU which is smaller than the actual-mtu " +
				"of the bridge (set by the mtu property), then manually set value will be ignored and the bridge " +
				"will act as if mtu=auto is set.",
		},
		"multicast_querier": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"multicast_router": {
			Type:     schema.TypeString,
			Optional: true,
		},
		KeyName: PropNameRw,
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0x8000",
			Description: "Bridge priority, used by STP to determine root bridge, used by MSTP to determine CIST " +
				"and IST regional root bridge. This property has no effect when protocol-mode is set to none.",
		},
		"protocol_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "rstp",
			Description: "Select Spanning tree protocol (STP) or Rapid spanning tree protocol (RSTP) to ensure a " +
				"loop-free topology for any bridged LAN.",
			ValidateFunc: validation.StringInSlice([]string{"none", "rstp", "stp", "mstp"}, false),
		},
		"pvid": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
			Description: "Port VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. " +
				"It applies e.g. to frames sent from bridge IP and destined to a bridge port. " +
				"This property only has effect when vlan-filtering is set to yes.",
			ValidateFunc: validation.IntBetween(1, 4094),
		},
		KeyRunning: PropRunningRo,
		"querier_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"query_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"query_response_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"startup_query_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"startup_query_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"transmit_hold_count": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      6,
			Description:  "The Transmit Hold Count used by the Port Transmit state machine to limit transmission rate.",
			ValidateFunc: validation.IntBetween(1, 10),
		},
		"vlan_filtering": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Globally enables or disables VLAN functionality for bridge.",
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

		Schema: resSchema,
	}
}
