package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceIPPool https://help.mikrotik.com/docs/display/ROS/IP+Pools
func ResourceIPPool() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/pool"),
		MetaId:           PropId(Name),

		KeyComment: PropCommentRw,
		KeyName:    PropNameForceNewRw,
		"next_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When address is acquired from pool that has no free addresses, and next-pool property is set " +
				"to another pool, then next IP address will be acquired from next-pool.",
		},
		"ranges": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: `IP address list of non-overlapping IP address ranges in form of: ` +
				`["from1-to1", "from2-to2", ..., "fromN-toN"]. ` +
				`For example, ["10.0.0.1-10.0.0.27", "10.0.0.32-10.0.0.47"]`,
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
