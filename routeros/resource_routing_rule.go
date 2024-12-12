package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    ".nextid": "*FFFFFFFF",
    "action": "lookup",
    "disabled": "false",
    "dst-address": "2.2.2.0/24",
    "inactive": "false",
    "interface": "bridge1",
    "routing-mark": "main",
    "src-address": "1.1.1.1/32",
    "table": "main"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Policy+Routing
func ResourceRoutingRule() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rule"),
		MetaId:           PropId(Id),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "An action to take on the matching packet:\n  * drop - silently drop the packet.\n  * lookup - perform a " +
				"lookup in routing tables.\n  * lookup-only-in-table - perform lookup only in the specified routing table " +
				"(see table parameter).\n  * unreachable - generate ICMP unreachable message and send it back to the source.",
			ValidateFunc:     validation.StringInSlice([]string{"drop", "lookup", "lookup-only-in-table", "unreachable"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"dst_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The destination address of the packet to match.",
		},
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Incoming interface to match.",
		},
		"min_prefix": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Equivalent to Linux IP rule `suppress_prefixlength`. For example to suppress the default route " +
				"in the routing decision set the value to 0.",
		},
		"routing_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Match specific routing mark.",
		},
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The source address of the packet to match.",
		},
		"table": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the routing table to use for lookup.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
