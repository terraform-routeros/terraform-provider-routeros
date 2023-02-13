package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCapsManProvisioning() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/provisioning"),
		MetaId:           PropId(Name),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"common_name_regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name_prefix": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hw_supported_modes": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_address_ranges": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"identity_regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"master_configuration": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name_format": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"radio_mac": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"slave_configurations": {
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
