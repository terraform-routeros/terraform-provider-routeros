package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "from-time": "0s",
    "limitation": "test",
    "profile": "test",
    "till-time": "23h59m59s",
    "weekdays": "sunday,monday,tuesday,wednesday,thursday,friday,saturday"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-ProfileLimitations
func ResourceUserManagerProfileLimitation() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/profile-limitation"),
		MetaId:           PropId(Id),

		"from_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "Time of the day when the limitation should take place.",
			DiffSuppressFunc: TimeEquall,
		},
		"limitation": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The limitation name.",
		},
		"profile": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The profile name.",
		},
		"till_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "23h59m59s",
			Description:      "Time of the day when the limitation should end.",
			DiffSuppressFunc: TimeEquall,
		},
		"weekdays": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}, false),
			},
			Description:      "Days of the week when the limitation is active.",
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
