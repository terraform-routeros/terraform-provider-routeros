package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "default": "true",
  "disabled": "false",
  "interface": "sfp1",
  "leds": "sfp-led",
  "type": "interface-activity"
}
*/

// https://help.mikrotik.com/docs/display/ROS/LEDs
func ResourceSystemLed() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/leds"),
		MetaId:           PropId(Id),

		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set the interface to which the LED is connected.",
			ValidateFunc:     validation.StringIsNotWhiteSpace,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"leds": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "An option to set the LED name.",
		},
		"modem_signal_treshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "An option to set the signal strength threshold for the modem LED.",
			ValidateFunc:     validation.IntBetween(-113, -51),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An option to set the LED type.",
			ValidateFunc: validation.StringInSlice([]string{
				"align-down", "align-left", "align-right", "align-up", "ap-cap", "fan-fault", "flash-access",
				"interface-activity", "interface-receive", "interface-speed", "interface-speed-1G",
				"interface-speed-25G", "interface-status", "interface-transmit",
				"modem-signal", "modem-technology", "off", "on", "poe-fault", "poe-out",
				"wireless-signal-strength", "wireless-status",
			}, false),
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
