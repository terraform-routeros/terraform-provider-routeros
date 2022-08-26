package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceSystemScheduler https://wiki.mikrotik.com/wiki/Manual:System/Scheduler
func ResourceSystemScheduler() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/scheduler"),
		MetaId:           PropId(Name),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Interval between two script executions, if time interval is set to zero, the script is only " +
				"executed at its start time, otherwise it is executed repeatedly at the time interval is specified.",
		},
		KeyName: PropNameRw,
		"next_run": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"on_event": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the script to execute. It must be presented at /system script.",
		},
		"owner": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"policy": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ftp", "reboot", "read", "write", "policy", "test",
					"password", "sniff", "sensitive", "romon", "dude"}, false),
			},
		},
		"run_count": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "This counter is incremented each time the script is executed.",
		},
		"start_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "Date of the first script execution.",
		},
		"start_time": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
			Description: "Time of the first script execution. If scheduler item has start-time set to startup, it " +
				"behaves as if start-time and start-date were set to time 3 seconds after console starts up. " +
				"It means that all scripts having start-time is startup and interval is 0 will be executed once each " +
				"time router boots. If the interval is set to value other than 0 scheduler will not run at startup.",
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
