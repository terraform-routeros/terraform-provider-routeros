package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceIPRoute https://wiki.mikrotik.com/wiki/Manual:IP/Route
func ResourceIPRoute() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/route"),
		MetaId:           PropId(Id),

		"active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag indicates whether the route is elected as Active and eligible to be added to the FIB.",
		},
		// If the parameter is present in the request, the MT automatically marks the route as a blackhole.
		// To solve this problem, let's remove the default value and the parameter will be present in the query
		// only when explicitly specified in the configuration.
		"blackhole": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "It's a blackhole route. If you need to cancel route marking, then simply delete the " +
				"parameter from the configuration of the TF. The value of the parameter (true or false) has no " +
				"effect on the MT processing logic.",
		},
		"check_gateway": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Currently used check-gateway option.",
			ValidateFunc: validation.StringInSlice([]string{"arp", "bfd", "bfd-multihop", "none", "ping"}, false),
		},
		KeyComment: PropCommentRw,
		"dhcp": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag indicates whether the route was added by the DHCP service.",
		},
		KeyDisabled: PropDisabledRw,
		"distance": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			Description:  "Value used in route selection. Routes with smaller distance value are given preference.",
			ValidateFunc: validation.IntBetween(1, 255),
		},
		"dst_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0.0.0.0/0", // Without the default setting, a non-empty plan is returned.
			Description:  "IP prefix of route, specifies destination addresses that this route can be used for.",
			ValidateFunc: validation.IsCIDR,
		},
		KeyDynamic: PropDynamicRo,
		"ecmp": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag indicates whether the route is added as an Equal-Cost Multi-Path route in the FIB.",
		},
		"gateway": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Array of IP addresses or interface names. Specifies which host or interface packets should " +
				"be sent to (IP | interface | IP%interface | IP@table[, IP | string, [..]]).",
		},
		"hw_offloaded": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether the route is eligible to be hardware offloaded on supported hardware.",
		},
		"immediate_gw": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows actual (resolved) gateway and interface that will be used for packet forwarding.",
		},
		KeyInactive: PropInactiveRo,
		"pref_src": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Which of the local IP addresses to use for locally originated packets that are sent via this " +
				"route. Value of this property has no effect on forwarded packets. If value of this property is set " +
				"to IP address that is not local address of this router then the route will be inactive (in ROS v6, " +
				"ROS v7 allows IP spoofing).",
			ValidateFunc: validation.IsIPv4Address,
		},
		"routing_table": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "main",
			Description: "Routing table this route belongs to.",
		},
		"scope": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
			Description: "Used in nexthop resolution. Route can resolve nexthop only through routes that have scope " +
				"less than or equal to the target-scope of this route.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"static": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"suppress_hw_offload": {
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"target_scope": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
			Description: "Used in nexthop resolution. This is the maximum value of scope for a route through which a " +
				"nexthop of this route can be resolved.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"vrf_interface": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "VRF interface name.",
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
