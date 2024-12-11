package routeros

import (
	"regexp"

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
	var reIPv4Range = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s*-\s*(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/address-list"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "A single IP address or range of IPs to add to address list or DNS name. You can input for " +
				"example, '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on " +
				"saving.",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				ips := reIPv4Range.FindStringSubmatch(new)
				if len(ips) == 3 {
					s, _ := IpRangeToCIDR(ips[1], ips[2])
					return old == s
				}

				return false
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
			Description: "Name for the address list of the added IP address.",
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
