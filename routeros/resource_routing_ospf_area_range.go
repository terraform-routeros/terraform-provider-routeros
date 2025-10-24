package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "advertise": "true",
    "area": "ospf-area-1",
    "cost": "100",
    "disabled": "false",
    "inactive": "false",
    "prefix": "::/0"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/331612216/routing+ospf#id-/routing/ospf-/routing/ospf/area/range
func ResourceRoutingOspfAreaRange() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/ospf/area/range"),
		MetaId:           PropId(Id),

		"advertise": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to create a summary LSA and advertise it to the adjacent areas.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"area": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OSPF area associated with this range.",
		},
		KeyComment: PropCommentRw,
		"cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The cost of the summary LSA this range will createdefault - use the largest cost of all routes " +
				"used (i.e. routes that fall within this range).",
		},
		KeyDisabled: PropDisabledRw,
		"prefix": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The network prefix of this range.",
			ValidateFunc: validation.IsCIDR,
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
