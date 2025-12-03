package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "broadcast": "false",
    "comment": "Something",
    "disabled": "false",
    "inactive": "true",
    "instance": "zt1",
    "ip-range": "192.168.88.200-192.168.88.220",
    "ip6-6plane": "false",
    "ip6-range": "fd00:feed:feed:beef::-fd00:feed:feed:beef:ffff:ffff:ffff:ffff",
    "ip6-rfc4193": "false",
    "mtu": "1598",
    "multicast-limit": "32",
    "name": "test",
    "network": "1234567812345678",
    "private": "true",
    "routes": "192.168.88.0/24@192.168.88.1,0.0.0.0/0@192.168.88.1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/ZeroTier#ZeroTier-Parameters.1
func ResourceZerotierController() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/zerotier/controller"),
		MetaId:           PropId(Id),

		"broadcast": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "An option to allow receiving broadcast packets.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"inactive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the ZeroTier network is inactive.",
		},
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier instance name.",
		},
		"ip_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The IP range of the ZeroTier network.",
		},
		"ip6_6plane": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option to assign every member a `/80` address within a `/40` network with using NDP emulation.",
		},
		"ip6_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The IPv6 range of the ZeroTier network.",
		},
		"ip6_rfc4193": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option to assign every member a `/128` address within a `/88` network.",
		},
		KeyMtu: PropL2MtuRw,
		"multicast_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     32,
			Description: "An option to limit the maximum recipients of a multicast packet.",
		},
		KeyName: PropName("Name of the ZeroTier controller."),
		"network": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier network identifier.",
		},
		"private": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "The ZeroTier network access control.",
		},
		"routes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "The routes list that will be pushed to the client.",
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
