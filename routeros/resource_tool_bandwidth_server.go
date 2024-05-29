package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "allocate-udp-ports-from": "2000",
  "authenticate": "true",
  "enabled": "true",
  "max-sessions": "100"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Bandwidth+Test
func ResourceToolBandwidthServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/bandwidth-server"),
		MetaId:           PropId(Id),

		"allocate_udp_ports_from": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Beginning of UDP port range.",
			ValidateFunc:     validation.IntBetween(1000, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"authenticate": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Communicate only with authenticated clients.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyEnabled: PropEnabled("Defines whether bandwidth server is enabled or not."),
		"max_sessions": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximal simultaneous test count.",
			ValidateFunc:     validation.IntBetween(1, 1000),
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
