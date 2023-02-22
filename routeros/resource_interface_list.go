package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/list"),
		MetaId:           PropId(Name),

		"builtin": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"comment": PropCommentRw,
		"dynamic": PropDynamicRo,
		"exclude": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"include": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": PropNameRw,
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
