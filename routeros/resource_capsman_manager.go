package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManManager() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/manager"),
		MetaId:           PropId(Name),
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"upgrade_policy": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
		},
		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
		},
		"ca_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
		},
		"require_peer_certificate": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"package_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
