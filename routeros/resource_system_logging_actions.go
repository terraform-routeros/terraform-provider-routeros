package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*3",
    "bsd-syslog": "false",
    "default": "true",
    "name": "remote",
    "remote": "0.0.0.0",
    "remote-port": "514",
    "src-address": "0.0.0.0",
    "syslog-facility": "daemon",
    "syslog-severity": "auto",
    "syslog-time-format": "bsd-syslog",
    "target": "remote"
  }
*/

// ResourceSystemLoggingActions https://help.mikrotik.com/docs/display/ROS/Log#Log-Actions
func ResourceSystemLoggingAction() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/logging/action"),
		MetaId:           PropId(Id),

		"bsd_syslog": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: `Whether to use bsd-syslog as defined in RFC 3164.`,
		},
		KeyDefault: PropDefaultRo,
		"disk_file_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Specifies number of files used to store log messages, applicable only if `action=disk`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"disk_file_name": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the file used to store log messages, applicable only if `action=disk`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"disk_lines_per_file": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Specifies maximum size of file in lines, applicable only if `action=disk`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"disk_stop_on_full": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to stop to save log messages to disk after the specified disk-lines-per-file " +
				"and disk-file-count number is reached, applicable only if `action=disk`.",
		},
		"email_start_tls": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to use tls when sending email, applicable only if `action=email`.",
		},
		"email_to": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Email address where logs are sent, applicable only if `action=email`.",
		},
		"memory_lines": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of records in local memory buffer, applicable only if `action=memory`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"memory_stop_on_full": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to stop to save log messages in local buffer after the specified memory-lines " +
				"number is reached.",
		},
		KeyName: PropName("Name of an action."),
		"remember": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to keep log messages, which have not yet been displayed in console, applicable " +
				"if `action=echo`.",
		},
		"remote": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Remote logging server's IP/IPv6 address, applicable if `action=remote`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Remote logging server's UDP port, applicable if `action=remote`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Source address used when sending packets to remote server, applicable if `action=remote`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"syslog_facility": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SYSLOG facility, applicable if `action=remote`.",
			ValidateFunc: validation.StringInSlice([]string{"auth", "authpriv", "cron", "daemon", "ftp",
				"kern", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7", "lpr",
				"mail", "news", "ntp", "syslog", "user", "uucp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"syslog_severity": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Severity level indicator defined in RFC 3164, applicable if `action=remote`.",
			ValidateFunc: validation.StringInSlice([]string{"alert", "auto", "critical", "debug", "emergency",
				"error", "info", "notice", "warning"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"syslog_time_format": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SYSLOG time format (`bsd-syslog` or `iso8601`).",
			ValidateFunc:     validation.StringInSlice([]string{"bsd-syslog", "iso8601"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"target": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Storage facility or target of log messages.",
			ValidateFunc:     validation.StringInSlice([]string{"disk", "echo", "email", "memory", "remote"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		switch d.Get("name").(string) {
		case "disk", "echo", "memory", "remote":
			d.SetId(d.Get("name").(string))
			id, err := dynamicIdLookup(Name, resSchema[MetaResourcePath].Default.(string), m.(Client), d)
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
				return diag.FromErr(err)
			}
			d.SetId(id)

			return ResourceUpdate(ctx, resSchema, d, m)
		default:
			return ResourceCreate(ctx, resSchema, d, m)
		}
	}

	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		switch d.Get("name").(string) {
		case "disk", "echo", "memory", "remote":
			return ResourceRead(ctx, resSchema, d, m)
		default:
			return ResourceRead(ctx, resSchema, d, m)
		}
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		switch d.Get("name").(string) {
		case "disk", "echo", "memory", "remote":
			return SystemResourceDelete(ctx, resSchema, d, m)
		default:
			return ResourceDelete(ctx, resSchema, d, m)
		}
	}

	return &schema.Resource{
		CreateContext: resCreate,
		ReadContext:   resRead,
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
