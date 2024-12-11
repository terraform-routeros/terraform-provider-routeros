package routeros

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
		"address": "123:dead::beaf/64",
    "creation-time": "sep/29/2022 07:09:24",
    "disabled": "false",
    "dynamic": "true",
    "list": "AAA",
    "timeout": "0s"
  }
*/

// ResourceIPv6FirewallAddrList https://help.mikrotik.com/docs/display/ROS/Address-lists
// They work more or less the same as IPv4 address lists, except no ranges
func ResourceIPv6FirewallAddrList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/firewall/address-list"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A single IPv6 address or IPv6 CIDR subnet",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				return old == fmt.Sprintf("%s/128", new)
			},
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
			Description: "Name for the address list of the added IPv6 address.",
		},
		"timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Time after address will be removed from address list. If timeout is not specified,
the address will be stored into the address list permanently.  
	> Please plan your work logic based on the fact that after the timeout    
	> the resource has been destroyed outside of a Terraform. 
`,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				// Compare intervals:
				oDuration, err := ParseDuration(old)
				if err != nil {
					panic("[FirewallAddrList Timeout] parse 'old' duration error: " + err.Error())
				}

				nDuration, err := ParseDuration(new)
				if err != nil {
					panic("[FirewallAddrList Timeout] parse 'new' duration error: " + err.Error())
				}

				//                     old       new
				// ~ timeout       = "4m59s" -> "5m"
				return nDuration.Seconds() > oDuration.Seconds()
			},
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
