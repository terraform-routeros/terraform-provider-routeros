package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceIPRoute https://wiki.mikrotik.com/wiki/Manual:Virtual_Routing_and_Forwarding
func ResourceIPVrf() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/vrf"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		KeyComment:  PropCommentRw,
		KeyName:     PropName("Unique name of the VRF."),
		"interfaces": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "At least one interface must be added to the VRF.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MinItems: 1,
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
