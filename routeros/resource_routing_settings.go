package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "check-gateway-ping-count": "2",
  "check-gateway-ping-interval": "10s",
  "check-gateway-ping-timeout": "1s",
  "policy-rules": "mangle,vrf-lookup,vrf-unreach,local,user,main",
  "single-process": "false"
}
*/

// ResourceRoutingSettings manages the /routing/settings menu. The policy-rules
// and single-process properties are available since RouterOS 7.22.
// https://help.mikrotik.com/docs/spaces/ROS/pages/179487340/Routing
func ResourceRoutingSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/settings"),
		MetaId:           PropId(Id),

		"check_gateway_ping_count": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "How many ICMP ping requests will be sent to the gateway.",
		},
		"check_gateway_ping_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Description:      "Interval between ICMP ping requests sent to the gateway.",
			DiffSuppressFunc: TimeEqual,
		},
		"check_gateway_ping_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Description:      "Timeout for ICMP ping requests sent to the gateway.",
			DiffSuppressFunc: TimeEqual,
		},
		"policy_rules": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Description:      "Comma separated list of routing policy rules that are applied to the packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"single_process": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable single-process routing protocol management.",
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