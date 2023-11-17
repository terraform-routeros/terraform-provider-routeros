package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "name": "test",
    "name-for-users": "test",
    "override-shared-users": "off",
    "price": "0",
    "starts-when": "assigned",
    "validity": "unlimited"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Profiles
func ResourceUserManagerProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/profile"),
		MetaId:           PropId(Id),

		KeyName: PropName("Unique name of the profile."),
		"name_for_users": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The name that will be shown to users in the web interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"override_shared_users": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "off",
			Description: "An option whether to allow multiple sessions with the same user name.",
		},
		"price": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Default:      .0,
			Description:  "The price of the profile.",
			ValidateFunc: validation.FloatAtLeast(.0),
		},
		"starts_when": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "assigned",
			Description:  "The time when the profile becomes active (`assigned` - immediately when the profile entry is created, `first-auth` - upon first authentication request).",
			ValidateFunc: validation.StringInSlice([]string{"assigned", "first-auth"}, false),
		},
		"validity": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "unlimited",
			Description: "The total amount of time a user can use this profile.",
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
