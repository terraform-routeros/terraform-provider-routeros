package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address": "0.0.0.0/0",
    "comment": "system default user",
    "disabled": "false",
    "expired": "true",
    "group": "full",
    "last-logged-in": "may/02/2023 05:21:45",
    "name": "admin"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/User#User-RouterUsers
func ResourceUser() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("last_logged_in"),

		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Host or network address from which the user is allowed to log in.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"expired": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Password expired.",
		},
		"group": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the group the user belongs to.",
		},
		"inactivity_policy": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Inactivity policy.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "lockscreen", "logout"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"inactivity_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Inactivity timeout for non-GUI sessions.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName("User name. Although it must start with an alphanumeric character, it may contain '*', " +
			"'_', '.' and '@' symbols."),
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "User  password. If not specified, it is left blank (hit [Enter] when logging  in). It " +
				"conforms to standard Unix characteristics of passwords and may  contain letters, digits, " +
				"'*' and '_' symbols.",
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
