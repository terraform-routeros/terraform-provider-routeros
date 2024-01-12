package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*2",
    "chain": "123",
    "comment": "test",
    "disabled": "true",
    "inactive": "true",
    "rule": "if (active) {accept}"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Route+Selection+and+Filters
func ResourceRoutingFilterRule() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/filter/rule"),
		MetaId:           PropId(Id),

		"chain": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Chain name.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"rule": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Filter rule.",
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
