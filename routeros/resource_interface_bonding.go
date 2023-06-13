package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceBonding https://wiki.mikrotik.com/wiki/Manual:Interface/Bonding
func ResourceInterfaceBonding() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bonding"),
		MetaId:           PropId(Name),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyName:     PropNameForceNewRw,
		KeyArp:      PropArpRw,
		"mtu": {
			Type:        schema.TypeString,
			Default:     "1500",
			Optional:    true,
			Description: "Maximum Transmit Unit in bytes. Must be smaller or equal to the smallest L2MTU value of a bonding slave.",
		},
		"arp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			// Default:  "100ms",
			Computed:         true,
			Description:      "Time in milliseconds which defines how often to monitor ARP requests",
			DiffSuppressFunc: TimeEquall,
			ValidateFunc:     ValidationTime,
		},
		"arp_ip_targets": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "IP target address which will be monitored if link-monitoring is set to " +
				"arp. You can specify multiple IP addresses.",
		},
		"down_delay": {
			Type:     schema.TypeString,
			Optional: true,
			// Default:  "0s",
			Computed: true,
			Description: "If a link failure has been detected, bonding interface is disabled for " +
				"down-delay time. Value should be a multiple of mii-interval, otherwise it will be rounded down to the nearest value.",
			DiffSuppressFunc: TimeEquall,
			ValidateFunc:     ValidationTime,
		},
		"forced_mac_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "By default, bonding interface will use MAC address of the first selected slave interface. " +
				"This property allows to configure static MAC address for the bond interface.",
		},
		"lacp_rate": {
			Type:     schema.TypeString,
			Optional: true,
			// Default:  "2m5s",
			Computed: true,
			Description: "Link Aggregation Control Protocol rate specifies how often to exchange with LACPDUs between bonding peer. " +
				"Used to determine whether link is up or other changes have occurred in the network. LACP tries to adapt to these changes providing failover.",
			DiffSuppressFunc: TimeEquall,
			ValidateFunc:     ValidationTime,
		},
		"link_monitoring": {
			Type:         schema.TypeString,
			Default:      "mii",
			Optional:     true,
			Description:  "Method to use for monitoring the link (whether it is up or down)",
			ValidateFunc: validation.StringInSlice([]string{"arp", "mii", "none"}, true),
		},
		"min_links": {
			Type: schema.TypeInt,
			//Default:      0,
			Optional:     true,
			Description:  "How many active slave links needed for bonding to become active",
			ValidateFunc: validation.IntBetween(0, 4294967295),
		},
		"mii_interval": {
			Type:     schema.TypeString,
			Optional: true,
			// Default:  "100ms",
			Computed:         true,
			Description:      "How often to monitor the link for failures (parameter used only if link-monitoring is mii)",
			DiffSuppressFunc: TimeEquall,
			ValidateFunc:     ValidationTime,
		},
		"mode": {
			Type:        schema.TypeString,
			Default:     "balance-rr",
			Optional:    true,
			Description: "Specifies one of the bonding policies",
			ValidateFunc: validation.StringInSlice([]string{"802.3ad", "active-backup", "balance-alb",
				"balance-rr", "balance-tlb", "balance-xor", "broadcast"}, true),
		},
		"primary": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Controls the primary interface between active slave ports, works only for active-backup, balance-tlb and balance-alb modes.",
		},
		"slaves": {
			Type:     schema.TypeString,
			Required: true,
			//Default:     "none",
			Description: "At least two ethernet-like interfaces separated by a comma, which will be used for bonding",
		},
		"up_delay": {
			Type:     schema.TypeString,
			Optional: true,
			// Default:  "0s",
			Computed: true,
			Description: "If a link has been brought up, bonding interface is disabled for up-delay time and after this " +
				"time it is enabled. Value should be a multiple of mii-interval, otherwise it will be rounded down to the nearest value.",
			DiffSuppressFunc: TimeEquall,
			ValidateFunc:     ValidationTime,
		},
		"transmit_hash_policy": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "layer-2",
			Description:  "Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes",
			ValidateFunc: validation.StringInSlice([]string{"layer-2", "layer-2-and-3", "layer-3-and-4"}, true),
		},
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
