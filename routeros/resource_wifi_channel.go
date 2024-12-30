package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "band": "2ghz-n",
    "disabled": "false",
    "frequency": "2412",
    "name": "channel1",
    "secondary-frequency": "disabled",
    "skip-dfs-channels": "disabled",
    "width": "20mhz"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Channelproperties
func ResourceWifiChannel() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/channel"),
		MetaId:           PropId(Id),

		"band": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Frequency band and wireless standard that will be used by the AP. ",
			ValidateFunc: validation.StringInSlice([]string{"2ghz-g", "2ghz-n", "2ghz-ax", "5ghz-a", "5ghz-ac", "5ghz-ax",
				"5ghz-an", "5ghz-n"}, false),
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"frequency": {
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "Channel frequency value or range in MHz on which AP or station will operate.",
		},
		KeyName: PropName("Name of the channel."),
		"reselect_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option that specifies when the interface should rescan channel availability and select the most appropriate one to use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secondary_frequency": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
			Description: "Specifies the second frequency that will be used for 80+80MHz configuration. " +
				"Set it to `disabled` in order to disable 80+80MHz capability.",
		},
		"skip_dfs_channels": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to avoid using channels on which channel availability check (listening for the presence of radar signals) is required.",
			ValidateFunc: validation.StringInSlice([]string{"10min-cac", "all", "disabled"}, false),
		},
		"width": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Channel width.",
			ValidateFunc: validation.StringInSlice([]string{"20mhz", "20/40mhz", "20/40mhz-Ce", "20/40mhz-eC", "20/40/80mhz", "20/40/80+80mhz", "20/40/80/160mhz"}, false),
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
