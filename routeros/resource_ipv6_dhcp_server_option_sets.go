package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "name": "set1",
    "options": "dns,option1"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/
func ResourceIpv6DhcpServerOptionSets() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-server/option/sets"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		KeyName:    PropName("The name of the DHCPv6 option."),
		"options": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The list of options.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
