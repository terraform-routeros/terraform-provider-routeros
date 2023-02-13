package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManDatapath() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/datapath"),
		MetaId:           PropId(Name),

		"bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"bridge_cost": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		KeyName:    PropNameRw,
		KeyComment: PropCommentRw,
		"bridge_horizon": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"interface_list": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		KeyL2Mtu: PropL2MtuRo,
		"local_forwarding": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"client_to_client_forwarding": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"openflow_switch": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vlan_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vlan_id": {
			Type:     schema.TypeInt,
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
