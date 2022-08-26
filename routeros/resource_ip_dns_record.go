package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
		KeyDynamic:  PropDynamicRo,
		KeyDisabled: PropDisabledRw,
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the DNS hostname to be created.",
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
