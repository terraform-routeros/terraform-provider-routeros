package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "disabled": "false",
  "match-count": "0",
  "name-count": "0",
  "url": "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
}
*/

// ResourceDnsAdlist https://help.mikrotik.com/docs/display/ROS/DNS#DNS-Adlist
func ResourceDnsAdlist() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns/adlist"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("match_count", "name_count"),

		KeyDisabled: PropDisabledRw,
		"file": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Used to specify a local file path from which to read adlist data.",
			ExactlyOneOf: []string{"file", "url"},
		},
		"ssl_verify": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether to validate the server's SSL certificate when connecting to an online " +
				"resource. Will use the `/certificate` list to verify server validity.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"url": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Used to specify the URL of an adlist.",
			ExactlyOneOf: []string{"file", "url"},
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
