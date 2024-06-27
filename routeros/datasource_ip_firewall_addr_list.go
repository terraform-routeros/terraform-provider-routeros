package routeros

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func getIPFirewallAddrListSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				KeyFilter: PropFilterRw,
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyComment: {
					Type:     schema.TypeString,
					Computed: true,
				},
				"creation_time": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyDisabled: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				KeyDynamic: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"timeout": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}
