package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "attributes": "",
    "caller-id": "bind",
    "comment": "test",
    "disabled": "false",
    "group": "test",
    "name": "test",
    "otp-secret": "",
    "password": "password",
    "shared-users": "1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Users
func ResourceUserManagerUser() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/user"),
		MetaId:           PropId(Id),

		"attributes": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A custom set of colon-separated attributes with their values will be added to `Access-Accept` messages for users in this group.",
		},
		"caller_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Allow user's authentication with a specific Calling-Station-Id value.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"group": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the group the user is associated with.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Username for session authentication."),
		"otp_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A token of a one-time code that will be attached to the password.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The password of the user for session authentication.",
		},
		"shared_users": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The total amount of sessions the user can simultaneously establish.",
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
