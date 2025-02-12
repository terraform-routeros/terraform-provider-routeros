package routeros

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*FFFFFFFA",
    "default": "true",
    "kind": "pcq",
    "name": "pcq-upload-default",
    "pcq-burst-rate": "0",
    "pcq-burst-threshold": "0",
    "pcq-burst-time": "10s",
    "pcq-classifier": "src-address",
    "pcq-dst-address-mask": "32",
    "pcq-dst-address6-mask": "128",
    "pcq-limit": "50",
    "pcq-rate": "0",
    "pcq-src-address-mask": "32",
    "pcq-src-address6-mask": "128",
    "pcq-total-limit": "2000"
  },
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/328088/Queues#Queues-QueueTypes
// https://wiki.mikrotik.com/Manual:Queue#Queue_Types
func ResourceQueueType() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/type"),
		MetaId:           PropId(Id),

		KeyDefault: PropDefaultRo,
		"kind": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Queue kind.",
		},
		KeyName: PropName("Type name."),
		// BFIFO
		"bfifo_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of bytes that the BFIFO queue can hold. Applies if `kind` is `bfifo`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// MQ PFIFO
		"mq_pfifo_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Multi-queue PFIFO limit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// PFIFO
		"pfifo_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of packets that the PFIFO queue can hold. Applies if `kind` is `pfifo`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// RED
		"red_avg_packet": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Used by RED for average queue size calculations (for packet to byte translation).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"red_burst": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of packets allowed for bursts of packets when there are no packets in the queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"red_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "RED queue limit in packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"red_max_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The average queue size at which packet marking probability is the highest.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"red_min_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Average queue size in bytes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// SFQ
		"sfq_allot": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Amount of data in bytes that can be sent in one round-robin round.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sfq_perturb": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "How often hash function must be refreshed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// PCQ
		"pcq_burst_rate": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximal upload/download data rate which can be reached while the burst for substream is allowed.",
			DiffSuppressFunc: BitsEqual,
		},
		"pcq_burst_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "This is value of burst on/off switch.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_burst_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Period of time, in seconds, over which the average data rate is calculated. (This is " +
				"NOT the time of actual burst).",
			DiffSuppressFunc: TimeEqual,
		},
		"pcq_classifier": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Selection of sub-stream identifiers.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"src-address", "dst-address", "src-port", "dst-port"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_dst_address_mask": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Size of IPv4 network that will be used as dst-address sub-stream identifier.",
			ValidateFunc:     validation.IntBetween(0, 32),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_dst_address6_mask": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Size of IPV6 network that will be used as dst-address sub-stream identifier.",
			ValidateFunc:     validation.IntBetween(0, 128),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Queue size of a single sub-stream (in kilobytes).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_rate": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximal available data rate of each sub-steam.",
			DiffSuppressFunc: BitsEqual,
		},
		"pcq_src_address_mask": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Size of IPv4 network that will be used as src-address sub-stream identifier.",
			ValidateFunc:     validation.IntBetween(0, 32),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_src_address6_mask": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Size of IPV6 network that will be used as src-address sub-stream identifier.",
			ValidateFunc:     validation.IntBetween(0, 128),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcq_total_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Max amount of bytes queued (in kilobytes) for all sub-streams per PCQ instance. " +
				"Note that each queue tree entry has its own PCQ instance.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// CoDel
		"codel_ce_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Marks packets above a configured threshold with ECN.",
			DiffSuppressFunc: TimeEqualU(time.Nanosecond),
		},
		"codel_ecn": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option is used to mark packets instead of dropping them.",
		},
		"codel_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interval should be set on the order of the worst-case RTT through the bottleneck giving " +
				"endpoints sufficient time to react.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		"codel_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Queue limit, when the limit is reached, incoming packets are dropped.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"codel_target": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Represents an acceptable minimum persistent queue delay.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		// FQ-Codel
		"fq_codel_ce_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Marks packets above a configured threshold with ECN.",
			DiffSuppressFunc: TimeEqualU(time.Nanosecond),
		},
		"fq_codel_ecn": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option is used to mark packets instead of dropping them.",
		},
		"fq_codel_flows": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "A number of flows into which the incoming packets are classified.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fq_codel_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interval should be set on the order of the worst-case RTT through the bottleneck giving " +
				"endpoints sufficient time to react.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		"fq_codel_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Queue limit, when the limit is reached, incoming packets are dropped.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fq_codel_memlimit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "A total number of bytes that can be queued in this FQ-CoDel instance. Will be enforced " +
				"from the fq-codel-limit parameter.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fq_codel_quantum": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "A number of bytes used as 'deficit' in the fair queuing algorithm. Default (1514 bytes) " +
				"corresponds to the Ethernet MTU plus the hardware header length of 14 bytes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fq_codel_target": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Represents an acceptable minimum persistent queue delay.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		// CAKE
		"cake_ack_filter": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cake_atm": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Compensates for ATM cell framing, which is normally found on ADSL links.",
		},
		"cake_autorate_ingress": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Automatic capacity estimation based on traffic arriving at this qdisc. This is most likely " +
				"to be useful with cellular links, which tend to change quality randomly.  The Bandwidth Limit " +
				"parameter can be used in conjunction to specify an initial estimate. The shaper will periodically be " +
				"set to a bandwidth slightly below the estimated rate.  This estimator cannot estimate the bandwidth " +
				"of links downstream of itself.",
		},
		"cake_bandwidth": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Sets the shaper bandwidth.",
			DiffSuppressFunc: BitsEqual,
		},
		"cake_diffserv": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "CAKE can divide traffic into `tins` based on the Diffserv field:" +
				"\n  * `diffserv4` Provides a general-purpose Diffserv implementation with four tins: Bulk (CS1), " +
				"6.25% threshold, generally low priority. Best Effort (general), 100% threshold. Video (AF4x, AF3x, " +
				"CS3, AF2x, CS2, TOS4, TOS1), 50% threshold. Voice (CS7, CS6, EF, VA, CS5, CS4), 25% threshold." +
				"\n  * `diffserv3` (default) Provides a simple, general-purpose Diffserv implementation with three " +
				"tins: Bulk (CS1), 6.25% threshold, generally low priority. Best Effort (general), 100% threshold. " +
				"Voice (CS7, CS6, EF, VA, TOS4), 25% threshold, reduced Codel interval.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cake_flowmode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "\n  * `flowblind` - Disables flow isolation; all traffic passes through a single queue for each tin." +
				"\n  * `srchost` - Flows are defined only by source address." +
				"\n  * `dsthost` Flows are defined only by destination address." +
				"\n  * `hosts` - Flows are defined by source-destination host pairs. This is host isolation, rather " +
				"than flow isolation." +
				"\n  * `flows` - Flows are defined by the entire 5-tuple of source address, a destination address, " +
				"transport protocol, source port, and destination port. This is the type of flow isolation performed " +
				"by SFQ and fq_codel." +
				"\n  * `dual-srchost` Flows are defined by the 5-tuple, and fairness is applied first over source " +
				"addresses, then over individual flows. Good for use on egress traffic from a LAN to the internet, " +
				"where it'll prevent any LAN host from monopolizing the uplink, regardless of the number of flows they use." +
				"\n  * `dual-dsthost` Flows are defined by the 5-tuple, and fairness is applied first over destination " +
				"addresses, then over individual flows. Good for use on ingress traffic to a LAN from the internet, " +
				"where it'll prevent any LAN host from monopolizing the downlink, regardless of the number of flows they use." +
				"\n  * `triple-isolate` - Flows are defined by the 5-tuple, and fairness is applied over source *and* " +
				"destination addresses intelligently (ie. not merely by host-pairs), and also over individual flows." +
				"\n  * `nat` Instructs Cake to perform a NAT lookup before applying flow- isolation rules, to determine " +
				"the true addresses and port numbers of the packet, to improve fairness between hosts `inside` the NAT. " +
				"This has no practical effect in `flowblind` or `flows` modes, or if NAT is performed on a different host." +
				"\n  * `nonat` (default) The cake will not perform a NAT lookup. Flow isolation will be performed using " +
				"the addresses and port numbers directly visible to the interface Cake is attached to.",
			ValidateFunc: validation.StringInSlice([]string{"dsthost", "dual-dsthost", "dual-srchost", "flowblind",
				"flows", "hosts", "srchost", "triple-isolate"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cake_memlimit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Limit the memory consumed by Cake to LIMIT bytes. By default, the limit is calculated based " +
				"on the bandwidth and RTT settings.",
		},
		"cake_mpu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Rounds each packet (including overhead) up to a minimum length BYTES. ",
			ValidateFunc: validation.IntBetween(-64, 256),
		},
		"cake_nat": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Instructs Cake to perform a NAT lookup before applying a flow-isolation rule.",
		},
		"cake_overhead": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Adds BYTES to the size of each packet. BYTES may be negative.",
			ValidateFunc:     validation.IntBetween(-64, 256),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cake_overhead_scheme": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
		"cake_rtt": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Manually specify an RTT. Default 100ms is suitable for most Internet traffic.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		"cake_rtt_scheme": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "\n  * `datacentre` - For extremely high-performance 10GigE+ networks only. Equivalent to `RTT 100us`." +
				"\n  * `lan` - For pure Ethernet (not Wi-Fi) networks, at home or in the office. Don't use this when " +
				"shaping for an Internet access link. Equivalent to `RTT 1ms`." +
				"\n  * `metro` - For traffic mostly within a single city. Equivalent to `RTT 10ms`. regional For traffic " +
				"mostly within a European-sized country. Equivalent to `RTT 30ms`." +
				"\n  * `internet` (default) This is suitable for most Internet traffic. Equivalent to `RTT 100ms`." +
				"\n  * `oceanic` - For Internet traffic with generally above-average latency, such as that suffered by " +
				"Australasian residents. Equivalent to `RTT 300ms`." +
				"\n  * `satellite` - For traffic via geostationary satellites. Equivalent to `RTT 1000ms`." +
				"\n  * `interplanetary` - So named because Jupiter is about 1 light-hour from Earth. Use this to (almost) " +
				"completely disable AQM actions. Equivalent to `RTT 3600s`.",
			ValidateFunc: validation.StringInSlice([]string{"datacentre", "internet", "interplanetary", "lan", "metro",
				"none", "oceanic", "regional", "satellite"}, false),
		},
		"cake_wash": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Apply the wash option to clear all extra DiffServ (but not ECN bits), after priority queuing " +
				"has taken place.",
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
