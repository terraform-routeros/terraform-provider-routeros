package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    "accounting": "true",
    "interim-update": "0s",
    "use-circuit-id-in-nas-port-id": "false",
    "use-radius": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/PPP+AAA
func ResourcePppAaa() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ppp/aaa"),
		MetaId:           PropId(Id),

		"accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "An option that enables accounting for users.",
		},
		"enable_ipv6_accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option that enables IPv6 separate accounting.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interim_update": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "Interval between scheduled RADIUS Interim-Update messages.",
			DiffSuppressFunc: TimeEquall,
		},
		"use_circuit_id_in_nas_port_id": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"use_radius": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to use RADIUS server.",
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
