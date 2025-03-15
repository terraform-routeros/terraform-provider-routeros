package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "note": "For authorized use only.",
  "show-at-login": "true",
  "show-at-cli-login": "false"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/40992863/Note
func ResourceSystemNote() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/note"),
		MetaId:           PropId(Id),

		"note": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Note that will be displayed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"show_at_login": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to show system note on each login.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"show_at_cli_login": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to show system note before telnet login prompt.",
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
