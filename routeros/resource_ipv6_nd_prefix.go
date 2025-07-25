package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "6to4-interface": "none",
    "autonomous": "true",
    "disabled": "false",
    "interface": "lo",
    "invalid": "false",
    "on-link": "true",
    "preferred-lifetime": "1w",
    "prefix": "aaff::/64",
    "valid-lifetime": "4w2d"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/40992815/IPv6+Neighbor+Discovery#IPv6NeighborDiscovery-Prefix
func ResourceIpv6NdPrefix() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/nd/prefix"),
		MetaId:           PropId(Id),

		"6to4_interface": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If this option is specified, this prefix will be combined with the IPv4 address of the interface " +
				"name to produce a valid 6to4 prefix. The first 16 bits of this prefix will be replaced by 2002 and the " +
				"next 32 bits of this prefix will be replaced by the IPv4 address assigned to the interface name at configuration " +
				"time. The remaining 80 bits of the prefix (including the SLA ID) will be advertised as specified in " +
				"the configuration file.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"autonomous": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When set, indicates that this prefix can be used for autonomous address configuration. Otherwise, " +
				"prefix information is silently ignored.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface name on which stateless auto-configuration will be running.",
		},
		KeyInvalid: PropInvalidRo,
		"on_link": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When set, indicates that this prefix can be used for on-link determination. When not set the " +
				"advertisement makes no statement about the on-link or off-link properties of the prefix. For instance, " +
				"the prefix might be used for address configuration with some of the addresses belonging to the prefix " +
				"being on-link and others being off-link.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preferred_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Timeframe (relative to the time the packet is sent) after which generated address becomes " +
				"`deprecated`. Deprecated is used only for already existing connections and is usable until valid " +
				"lifetime expires.",
			DiffSuppressFunc: TimeEqual,
		},
		"prefix": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A prefix from which stateless address autoconfiguration generates the valid address.",
		},
		"valid_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The length of time (relative to the time the packet is sent) an address remains in the valid " +
				"state. The valid lifetime must be greater than or equal to the preferred lifetime.",
			DiffSuppressFunc: TimeEqual,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
