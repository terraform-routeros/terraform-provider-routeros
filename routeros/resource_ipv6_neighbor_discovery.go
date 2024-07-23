package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
[
  {
        ".id": "*3",
        "advertise-dns": "false",
        "advertise-mac-address": "true",
        "default": "false",
        "disabled": "false",
        "dns": "",
        "hop-limit": "unspecified",
        "interface": "vlan-wifi-XXX",
        "invalid": "false",
        "managed-address-configuration": "true",
        "mtu": "unspecified",
        "other-configuration": "true",
        "pref64": "",
        "ra-delay": "3s",
        "ra-interval": "3m20s-10m",
        "ra-lifetime": "30m",
        "ra-preference": "high",
        "reachable-time": "unspecified",
        "retransmit-interval": "unspecified"
    }
]
*/

// ResourceIPv6NeighborDiscovery https://help.mikrotik.com/docs/display/ROS/IPv6+Neighbor+Discovery
func ResourceIPv6NeighborDiscovery() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/nd"),
		MetaId:           PropId(Id),
		MetaDropByValue:  PropDropByValue("unspecified"),
		"advertise_dns": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Option to redistribute DNS server information using RADVD. You will need a running client-side software with Router Advertisement DNS support to take advantage of the advertised DNS information.",
			Default:     true,
		},
		"advertise_mac_address": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When set, the link-layer address of the outgoing interface is included in the RA.",
			Default:     true,
		},
		KeyComment: PropCommentRw,
		"default": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Neighbor discovery entry is the default configuration.",
		},
		"dns_servers": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify a single IPv6 address or list of addresses that will be provided to hosts for DNS server configuration.",
		},
		"dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify a single IPv6 address or comma separated list of addresses that will be provided to hosts for DNS server configuration.",
		},
		KeyDisabled: PropDisabledRw,
		"hop_limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The default value that should be placed in the Hop Count field of the IP header for outgoing (unicast) IP packets.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"interface": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The interface on which to run neighbor discovery." +
				"all - run ND on all running interfaces.",
		},
		KeyInvalid: PropInvalidRo,
		"managed_address_configuration": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Name of the IPv6 pool in which received IPv6 prefix will be added",
		},
		"mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The flag indicates whether hosts should use stateful autoconfiguration (DHCPv6) to obtain addresses",
			ValidateFunc: validation.IntBetween(0, 90000),
		},
		"other_configuration": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "The flag indicates whether hosts should use stateful autoconfiguration to obtain additional information (excluding addresses).",
		},
		"pref64": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Specify IPv6 prefix or list of prefixes within /32, /40. /48, /56, /64, or /96 subnet that will be provided to hosts as NAT64 prefixes.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ra_delay": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The minimum time allowed between sending multicast router advertisements from the interface.",
			Default:     "3s",
		},
		"ra_interval": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The min-max interval allowed between sending unsolicited multicast router advertisements from the interface.",
			Default:     "3m20s-10m",
		},
		"ra_preference": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify the router preference that is communicated to IPv6 hosts through router advertisements." +
				"The preference value in the router advertisements enables IPv6 hosts to select a default router to reach a remote destination",
			Default:      "medium",
			ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),
		},
		"ra_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify the router preference that is communicated to IPv6 hosts through router advertisements." +
				"The preference value in the router advertisements enables IPv6 hosts to select a default router to reach a remote destination",
			Default: "30m",
		},
		"reachable_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify the router preference that is communicated to IPv6 hosts through router advertisements." +
				"The preference value in the router advertisements enables IPv6 hosts to select a default router to reach a remote destination",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"retransmit_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time between retransmitted Neighbor Solicitation messages." +
				"Used by address resolution and the Neighbor Unreachability Detection algorithm (see Sections 7.2 and 7.3 of RFC 2461)",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
