package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/list"),
		MetaId:           PropId(Id),

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
		"name": PropNameForceNewRw,
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceInterfaceListV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
