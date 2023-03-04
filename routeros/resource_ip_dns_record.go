package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "address": "192.168.88.1",
    "comment": "defconf",
    "disabled": "false",
    "dynamic": "false",
    "name": "router.lan",
    "ttl": "1d"
  },
  {
    ".id": "*9",
    "address": "192.168.88.1",
    "disabled": "false",
    "dynamic": "false",
    "regexp": ".*pool.ntp.org",
    "ttl": "1d"
  },
*/

// ResourceDnsRecord https://wiki.mikrotik.com/wiki/Manual:IP/DNS
func ResourceDnsRecord() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns/static"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The A record to be returend from the DNS hostname.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The name of the DNS hostname to be created.",
			ExactlyOneOf: []string{"name", "regexp"},
		},
		"regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, " +
				"RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.",
			ExactlyOneOf: []string{"name", "regexp"},
		},
		"ttl": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ttl of the DNS record.",
		},
	}
	return &schema.Resource{
		Description: "Creates a DNS record on the MikroTik device.",

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
