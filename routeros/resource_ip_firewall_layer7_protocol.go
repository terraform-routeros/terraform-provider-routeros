package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "name": "rdp",
    "regexp": "rdpdr.*cliprdr.*rdpsnd"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/
func ResourceIpFirewallLayer7Protocol() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/layer7-protocol"),
		MetaId:           PropId(Id),

		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Descriptive name of l7 pattern used by configuration in firewall rules.",
		},
		"regexp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "POSIX compliant regular expression is used to match a pattern.",
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
