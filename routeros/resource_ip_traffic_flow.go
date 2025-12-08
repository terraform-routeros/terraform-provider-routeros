package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "active-flow-timeout": "30m",
  "cache-entries": "128k",
  "enabled": "false",
  "inactive-flow-timeout": "15s",
  "interfaces": "all",
  "packet-sampling": "false",
  "sampling-interval": "0",
  "sampling-space": "0"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/21102653/Traffic+Flow#TrafficFlow-General
func ResourceIpTrafficFlow() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/traffic-flow"),
		MetaId:           PropId(Id),

		"interfaces": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Names of those interfaces will be used to gather statistics for traffic-flow. To specify more " +
				"than one interface, separate them with a comma.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_entries": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Number of flows which can be in router's memory simultaneously.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"active_flow_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum life-time of a flow.",
			DiffSuppressFunc: TimeEqual,
		},
		"inactive_flow_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How long to keep the flow active, if it is idle. If a connection does not see any packet within " +
				"this timeout, then traffic-flow will send a packet out as a new flow. If this timeout is too small it " +
				"can create a significant amount of flows and overflow the buffer.",
			DiffSuppressFunc: TimeEqual,
		},
		"packet_sampling": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable or disable packet sampling feature.",
		},
		"sampling_interval": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The number of packets that are consecutively sampled.",
		},
		"sampling_space": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The number of packets that are consecutively omitted.",
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
