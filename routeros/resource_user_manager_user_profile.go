package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "end-time": "unlimited",
    "profile": "test",
    "state": "running",
    "user": "test"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-UserProfiles
func ResourceUserManagerUserProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/user-profile"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("end_time", "state"),

		"profile": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the profile to assign to the user.",
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the user to use the specified profile.",
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
