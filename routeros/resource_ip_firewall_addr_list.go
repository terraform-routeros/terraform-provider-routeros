package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "address": "0.0.0.0",
    "creation-time": "sep/29/2022 07:09:24",
    "disabled": "false",
    "dynamic": "true",
    "list": "AAA",
    "timeout": "0s"
  }
*/

// ResourceIPFirewallAddrList https://wiki.mikrotik.com/wiki/Manual:IP/Firewall/Address_list
func ResourceIPFirewallAddrList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/address-list"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "A single IP address or range of IPs to add to address list or DNS name. You can input for " +
				"example, '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on " +
				"saving.",
		},
		KeyComment: PropCommentRw,
		"creation_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Rule creation time",
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"list": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name for the address list of the added IP address.",
		},
		"timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time after address will be removed from address list. If timeout is not specified, " +
				"the address will be stored into the address list permanently.",
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
