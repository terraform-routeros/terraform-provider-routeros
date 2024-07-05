package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "all-leds-off": "never"
}
*/

// https://help.mikrotik.com/docs/display/ROS/LEDs#LEDs-LEDSettings
func ResourceSystemLedSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/leds/settings"),
		MetaId:           PropId(Id),

		"all_leds_off": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set when all LEDs should be turned off. Possible values: `after-1h`, `after-1min`, `immediate`, `never`.",
			ValidateFunc:     validation.StringInSlice([]string{"after-1h", "after-1min", "immediate", "never"}, false),
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
