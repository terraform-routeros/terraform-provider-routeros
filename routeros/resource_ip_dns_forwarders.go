package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "dns-servers": "1.1.1.1",
    "doh-servers": "2.2.2.2",
    "name": "aaa",
    "verify-doh-cert": "false"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/37748767/DNS#DNS-Forwarders
func ResourceIpDnsForwarders() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns/forwarders"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dns_servers": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "An IP address or DNS name of a domain name server. Can contain multiple records.",
		},
		"doh_servers": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A URL of DoH server. Can contain multiple records.",
		},
		KeyName: PropName("Forwarder name."),
		"verify_doh_cert": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether to validate the DoH server, when one is being used. Will use the `/certificate` " +
				"list in order to verify server validity.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
