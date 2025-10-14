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
		KeyComment: PropCommentRw,
		"deprioritize_unii_3_4": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to assign lower priority to channels with a control frequency of 5720 or 5825-5885 " +
				"MHz. These channels are unsupported by some client devices, making their automatic selection " +
				"undesirable. Defaults to `yes` in ETSI regulatory domains, elsewhere to `no`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
		"reselect_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the clock time when the interface should run \"rescan channel availability\" " +
				"and select the most appropriate one to use. Specifying the clock time will allow the system to select " +
				"this time dynamically and randomly. This helps to avoid a situation when many APs at the same time scan " +
				"the network, select the same channel, and prefer to use it at the same time. reselect-time uses a background " +
				"scan. " +
				"\nThe reselect process will choose the most suitable channel considering the number of networks in the " +
				"channel, channel usage, and overlap with networks in adjacent channels. It can be used with a list of " +
				"frequencies defined, or with frequency not set - using all supported frequencies." +
				"\nExample:\n" +
				"\n    - 01:00..01:30 → Would set the rescan of channels to run every night, once, randomly, between " +
				"01:00 AM to 01:30 AM, system clock time." +
				"\n    - 14:00..14:30 → Would set the rescan of channels to run every day (after midday), once, randomly " +
				"between 14:00:00 to 14:30:00 (or 2 PM to 2:30 PM), system clock time.",
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
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
