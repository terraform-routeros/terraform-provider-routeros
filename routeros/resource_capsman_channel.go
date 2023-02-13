package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManChannel() *schema.Resource {
	// return &schema.Resource{
	// 	Create: resourceCapsManChannelCreate,
	// 	Read:   resourceCapsManChannelRead,
	// 	Update: resourceCapsManChannelUpdate,
	// 	Delete: resourceCapsManChannelDelete,
	// 	Importer: &schema.ResourceImporter{
	// 		StateContext: schema.ImportStatePassthroughContext,
	// 	},

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/channel"),
		MetaId:           PropId(Name),

		"save_selected": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"width": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"control_channel_width": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		KeyName:    PropNameRw,
		KeyComment: PropCommentRw,
		"band": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"reselect_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"extension_channel": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"secondary_frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tx_power": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"skip_dfs_channels": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
