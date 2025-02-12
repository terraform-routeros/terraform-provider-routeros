package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    "accounting": "true",
    "default-group": "read",
    "exclude-groups": "full",
    "interim-update": "0s",
    "use-radius": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User#User-RemoteAAA
func ResourceUserAaa() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user/aaa"),
		MetaId:           PropId(Id),

		"accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "An option that enables accounting for users.",
		},
		"default_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "read",
			Description: "The user group that is used by default for users authenticated via a RADIUS server.",
		},
		"exclude_groups": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A set of groups that are not allowed for users authenticated by RADIUS.",
		},
		"interim_update": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "Interval between scheduled RADIUS Interim-Update messages.",
			DiffSuppressFunc: TimeEqual,
		},
		"use_radius": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to use RADIUS server.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
