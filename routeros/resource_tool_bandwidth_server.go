package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "allowed-interface-list": "LAN"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Bandwidth+Test
func ResourceToolBandwidthServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/bandwidth-server"),
		MetaId:           PropId(Id),

		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Defines whether bandwidth server is enabled or not.",
		},
		"authenticate": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Communicate only with authenticated clients.",
		},
		"max_sessions ": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     100,
			Description: "Maximal simultaneous test count.",
		},
		"allocate_udp_ports_from ": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     2000,
			Description: "Beginning of UDP port range.",
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
