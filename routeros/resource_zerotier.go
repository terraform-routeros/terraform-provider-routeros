package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "comment": "ZeroTier Central controller - https://my.zerotier.com/",
    "disabled": "false",
    "identity": "...",
    "identity.public": "...",
    "interfaces": "all",
    "name": "zt1",
    "online": "true",
    "port": "9993",
    "route-distance": "1",
    "state": "running"
}
*/

// https://help.mikrotik.com/docs/display/ROS/ZeroTier#ZeroTier-Parameters
func ResourceZerotier() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/zerotier"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("identity_public: identity.public"),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"identity": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The 40-bit unique instance address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"identity_public": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The public identity of the ZeroTier instance.",
		},
		"interfaces": {
			Type:             schema.TypeSet,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "The interfaces to discover ZeroTier peers by ARP and IP type connections.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the ZeroTier instance."),
		"online": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the ZeroTier instance is currently online.",
		},
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      9993,
			Description:  "The port number the instance listens to.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"route_distance": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The route distance for routes obtained from the planet/moon server.",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The state of the ZeroTier instance.",
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
