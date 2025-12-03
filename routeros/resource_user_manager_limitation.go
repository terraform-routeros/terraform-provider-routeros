package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "download-limit": "0",
    "name": "test",
    "rate-limit-burst-rx": "0",
    "rate-limit-burst-threshold-rx": "0",
    "rate-limit-burst-threshold-tx": "0",
    "rate-limit-burst-time-rx": "0s",
    "rate-limit-burst-time-tx": "0s",
    "rate-limit-burst-tx": "0",
    "rate-limit-min-rx": "0",
    "rate-limit-min-tx": "0",
    "rate-limit-priority": "0",
    "rate-limit-rx": "10",
    "rate-limit-tx": "10",
    "reset-counters-interval": "disabled",
    "reset-counters-start-time": "1970-01-01 00:00:00",
    "transfer-limit": "0",
    "upload-limit": "0",
    "uptime-limit": "0s"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Limitations
func ResourceUserManagerLimitation() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/limitation"),
		MetaId:           PropId(Id),

		"download_limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			Description:  "The total amount of traffic a user can download in bytes.",
			ValidateFunc: validation.IntAtLeast(0),
		},
		KeyName: PropName("Unique name of the limitation."),
		"rate_limit_burst_rx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_burst_tx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_burst_threshold_rx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_burst_threshold_tx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_burst_time_rx": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			DiffSuppressFunc: TimeEqual,
		},
		"rate_limit_burst_time_tx": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			DiffSuppressFunc: TimeEqual,
		},
		"rate_limit_min_rx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_min_tx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_rx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"rate_limit_tx": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		"reset_counters_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "disabled",
			Description:  "The interval from `reset_counters_start_time` when all associated user statistics are cleared.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "hourly", "daily", "weekly", "monthly"}, false),
		},
		"reset_counters_start_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "1970-01-01 00:00:00",
			Description: "Static date and time value from which `reset_counters_interval` is calculated.",
		},
		"transfer_limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			Description:  "The total amount of aggregated (download+upload) traffic in bytes.",
			ValidateFunc: validation.IntAtLeast(0),
		},
		"upload_limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			Description:  "The total amount of traffic a user can upload in bytes.",
			ValidateFunc: validation.IntAtLeast(0),
		},
		"uptime_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "The total amount of uptime a user can stay active.",
			DiffSuppressFunc: TimeEqual,
		},
	}

	return &schema.Resource{
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
