package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* {
  "date": "2024-05-17",
  "dst-active": "false",
  "gmt-offset": "+01:00",
  "time": "17:58:11",
  "time-zone-autodetect": "false",
  "time-zone-name": "Etc/GMT-1"
} */

func ResourceSystemClock() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/clock"),
		MetaId:           PropId(Id),
		"date": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Date.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: `This property has the value yes while daylight saving time of the current time zone is active.`,
		},
		"gmt_offset": {
			Type:     schema.TypeString,
			Computed: true,
			Description: `This is the current value of GMT offset used by the system, after applying base time zone ` +
				`offset and active daylight saving time offset.`,
		},
		"time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Time.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"time_zone_autodetect": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      `Feature available from v6.27. If enabled, the time zone will be set automatically.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"time_zone_name": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Name of the time zone. As most of the text values in RouterOS, this value is case ` +
				`sensitive. Special value manual applies ` +
				`[manually configured GMT offset](https://wiki.mikrotik.com/wiki/Manual:System/Time#Manual_time_zone_configuration), ` +
				`which by default is 00:00 with no daylight saving time.`,
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
