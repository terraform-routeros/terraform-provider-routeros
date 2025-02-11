package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "bucket-size": "0.1/0.1",
    "burst-limit": "0/0",
    "burst-threshold": "0/0",
    "burst-time": "0s/0s",
    "bytes": "0/0",
    "comment": "comment",
    "disabled": "false",
    "dropped": "0/0",
    "dynamic": "false",
    "invalid": "false",
    "limit-at": "0/0",
    "max-limit": "0/0",
    "name": "queue1",
    "packet-marks": "",
    "packet-rate": "0/0",
    "packets": "0/0",
    "parent": "none",
    "priority": "8/8",
    "queue": "default-small/default-small",
    "queued-bytes": "0/0",
    "queued-packets": "0/0",
    "rate": "0/0",
    "target": "1.0.1.0/24",
    "total-bytes": "0",
    "total-dropped": "0",
    "total-packet-rate": "0",
    "total-packets": "0",
    "total-queued-bytes": "0",
    "total-queued-packets": "0",
    "total-rate": "0"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/328088/Queues#Queues-SimpleQueue
// https://wiki.mikrotik.com/Manual:Queue#Simple_Queues
func ResourceQueueSimple() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/simple"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("bytes", "dropped", "packet_rate", "packets", "queued_bytes", "queued_packets",
			"rate", "total_bytes", "total_dropped", "total_packet_rate", "total_packets", "total_queued_bytes",
			"total_queued_packets", "total_rate"),

		"bucket_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"burst_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximal upload/download data rate which can be reached while the burst is active.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"burst_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When average data rate is below this value - burst is allowed, as soon as average data " +
				"rate reach this value - burst is denied (basically this is burst on/off switch). For optimal burst " +
				"behavior this value should above `limit-at` value and below `max-limit` value",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"burst_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Period of time, in seconds, over which the average upload/download data rate is calculated. " +
				"This is NOT the time of actual burst.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows to select only specific stream (from target address to this destination address) " +
				"for limitation explain what is target and what is dst and what is upload and what not.",
		},
		KeyDynamic: PropDynamicRo,
		KeyInvalid: PropInvalidRo,
		"limit_at": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Normal upload/download data rate that is guaranteed to a target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximal upload/download data rate that is allowed for a target to reach to reach what.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Queue name."),
		"packet_marks": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Allows to use marked packets from `/ip firewall mangle`. Take look at this packet flow diagram. " +
				"You need to make sure that packets are marked before the simple queues (before global-in HTB queue).",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"parent": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Assigns this queue as a child queue for selected target. Target queue can be HTB queue " +
				"or any other previously created queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Prioritize one child queue over other child queue. Does not work on parent queues (if queue " +
				"has at least one child). One is the highest, eight is the lowest priority. Child queue with higher " +
				"priority will have chance to reach its `max-limit` before child with lower priority. Priority have " +
				"nothing to do with bursts.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"queue": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Choose the type of the queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"target": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "List of IP address ranges that will be limited by this queue.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allow to specify time when particular queue will be active. Router must have correct time settings.",
			DiffSuppressFunc: TimeEqual,
		},
		"total_bucket_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_burst_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_burst_threshold": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_burst_time": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: TimeEqual,
		},
		"total_limit_at": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_max_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_priority": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
		},
		"total_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
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
