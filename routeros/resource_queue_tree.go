package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1000000",
    "bucket-size": "0.1",
    "burst-limit": "0",
    "burst-threshold": "0",
    "burst-time": "0s",
    "bytes": "0",
    "comment": "comment",
    "disabled": "false",
    "dropped": "0",
    "invalid": "false",
    "limit-at": "0",
    "max-limit": "0",
    "name": "queue1",
    "packet-mark": "",
    "packet-rate": "0",
    "packets": "0",
    "parent": "global",
    "priority": "8",
    "queue": "default-small",
    "queued-bytes": "0",
    "queued-packets": "0",
    "rate": "0"
  }

*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/328088/Queues#Queues-QueueTree
// https://wiki.mikrotik.com/Manual:Queue#Queue_Tree
func ResourceQueueTree() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/tree"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("borrows", "bytes", "dropped", "lends", "packet_rate", "packets", "pcq_queues",
			"queued_bytes", "queued_packets", "rate"),

		"bucket_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"burst_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximal data rate which can be reached while the burst is active.",
			DiffSuppressFunc: BitsEqual,
		},
		"burst_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When average data rate is below this value - burst is allowed, as soon as average data " +
				"rate reach this value - burst is denied (basically this is burst on/off switch). For optimal burst " +
				"behavior this value should above `limit-at` value and below `max-limit` value.",
			DiffSuppressFunc: BitsEqual,
		},
		"burst_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Period of time, in seconds, over which the average data rate is calculated. " +
				"This is NOT the time of actual burst.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInvalid:  PropInvalidRo,
		"limit_at": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Normal data rate that is guaranteed to a target.",
			DiffSuppressFunc: BitsEqual,
		},
		"max_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximal data rate that is allowed for a target to reach.",
			DiffSuppressFunc: BitsEqual,
		},
		KeyName: PropName("Queue tree name."),
		"packet_mark": {
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
			Required: true,
			Description: "Assigns this queue as a child queue for selected target. Target queue can be HTB queue " +
				"or any other previously created queue.",
		},
		"priority": {
			Type:     schema.TypeInt,
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
