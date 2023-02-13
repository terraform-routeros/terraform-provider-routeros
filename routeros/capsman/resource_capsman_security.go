package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCapsManSecurity() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/security"),
		MetaId:           PropId(Name),
		"group_encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"authentication_types": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"eap_methods": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"eap_radius_accounting": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		KeyName: PropNameRw,
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"passphrase": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"group_key_update": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tls_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tls_mode": {
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
