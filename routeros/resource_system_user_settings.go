package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    "minimum-categories": "0",
    "minimum-password-length": "0"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User#User-UserSettings
func ResourceSystemUserSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user/settings"),
		MetaId:           PropId(Id),

		"minimum_categories": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			Description:  "An option specifies the complexity requirements of the password, with categories being uppercase, lowercase, digit, and symbol.",
			ValidateFunc: validation.IntBetween(0, 4),
		},
		"minimum_password_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			Description:  "An option specifies the minimum length of the password.",
			ValidateFunc: validation.IntAtLeast(0),
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
