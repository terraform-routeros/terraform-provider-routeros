package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 {
    "active-ipv4": "yes",
    "active-ipv6": "yes",
    "enabled": "yes",
    "generic-timeout": "10m",
    "icmp-timeout": "10s",
    "loose-tcp-tracking": "true",
    "max-entries": "1048576",
    "tcp-close-timeout": "1m",
    "tcp-close-wait-timeout": "1m",
    "tcp-established-timeout": "1d",
    "tcp-fin-wait-timeout": "1m",
    "tcp-last-ack-timeout": "1m",
    "tcp-max-retrans-timeout": "5m",
    "tcp-syn-received-timeout": "5s",
    "tcp-syn-sent-timeout": "5s",
    "tcp-time-wait-timeout": "1m",
    "tcp-unacked-timeout": "5m",
    "total-entries": "87",
    "udp-stream-timeout": "3m",
    "udp-timeout": "10s"
}
*/

// ResourceIPConnectionTracking https://help.mikrotik.com/docs/display/ROS/Connection+tracking
func ResourceIPConnectionTracking() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/connection/tracking"),
		MetaId:           PropId(Name),
		MetaSkipFields:   PropSkipFields(`"total_entries"`),

		"active_ipv4": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "documentation is missing",
		},
		"active_ipv6": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "documentation is missing",
		},
		"enabled": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Allows to disable or enable connection tracking. Disabling connection tracking will cause several firewall features to stop working. 
				          See the list of affected features. Starting from v6.0rc2 default value is auto. This means that connection tracing is disabled until at least one firewall rule is added.`,
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"generic_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Timeout for all other connection entries",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"icmp_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "ICMP connection timeout",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"loose_tcp_tracking": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Disable picking up already established connections",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_entries": {
			Type: schema.TypeString,
			Description: `Max amount of entries that the connection tracking table can hold. This value depends on the installed amount of RAM.
                          Note that the system does not create a maximum_size connection tracking table when it starts, it may increase if the situation demands it and the system still has free ram, but size will not exceed 1048576`,
			Computed: true,
		},
		"tcp_close_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_close_wait_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_established_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time when established TCP connection times out.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_fin_wait_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_last_ack_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_max_retrans_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			// Documentation did contain the default, I'm getting it from the docker image default (7.10)
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_syn_received_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "TCP SYN timeout.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_syn_sent_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "TCP SYN timeout.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_time_wait_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_unacked_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "No documentation",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"udp_stream_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies the timeout of UDP connections that has seen packets in both directions",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"udp_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies the timeout for UDP connections that have seen packets in one direction",
			ValidateFunc:     ValidationTime,
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
