package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceListMember() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/list/member"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dynamic": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interface": {
			Type:     schema.TypeString,
			Required: true,
		},
		"list": {
			Type:     schema.TypeString,
			Required: true,
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
