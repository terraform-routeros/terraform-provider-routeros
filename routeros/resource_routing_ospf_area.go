package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*54",
    "name": "",
    "area-id": "",
    "default-cost": "",
    "instance": "",
    "no-summaries": "",
    "nssa-translate": "",
    "type": "",
  }
*/

// ResourceRoutingOspfArea https://help.mikrotik.com/docs/display/ROS/OSPF
func ResourceRoutingOspfArea() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/ospf/area"),
		MetaId:           PropId(Name),

		KeyName:     PropNameForceNewRw,
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"area_id": {
			Type:        schema.TypeString,
			Default:     "0.0.0.0",
			Optional:    true,
			Description: "OSPF area identifier.",
		},
		"default_cost": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Default cost of injected LSAs into the area.",
		},
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the OSPF instance this area belongs to.",
		},
		"no_summaries": {
			Type:        schema.TypeBool,
			Default:     false,
			Optional:    true,
			Description: "If set then the area will not flood summary LSAs in the stub area.",
		},
		"nssa_translate": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The parameter indicates which ABR will be used as a translator from type7 to type5 LSA.",
			ValidateFunc: validation.StringInSlice([]string{"no", "yes", "candidate"}, false),
		},
		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "default",
			Description:  "The area type.",
			ValidateFunc: validation.StringInSlice([]string{"default", "nssa", "stub"}, true),
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
