package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "enabled": "false",
  "hold-time": "0s..1m",
  "on-event": ""
}
*/

// https://help.mikrotik.com/docs/display/ROS/RouterBOARD#RouterBOARD-ModeandResetbuttons
func ResourceSystemRouterboardButtonMode() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/routerboard/mode-button"),
		MetaId:           PropId(Id),

		KeyEnabled: PropEnabled("An option to enable the operation of the button."),
		"hold_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to define the period within which the button should be pressed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"on_event": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set the script that will be run upon pressing the button.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
