package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address": "0.0.0.1",
    "comment": "comment",
    "disabled": "false",
    "mac-address": "00:00:00:00:01:10",
    "to-address": "0.0.0.2"
  }
*/

// https://help.mikrotik.com/docs/pages/viewpage.action?pageId=56459266#HotSpot(Captiveportal)-IPBinding
func ResourceIpHotspotIpBinding() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/ip-binding"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The original IP address of the client.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "MAC address of the client.",
		},
		"server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the HotSpot server. `all` - will be applied to all hotspot servers.",
		},
		"to_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "New IP address of the client, translation occurs on the router (client does not know anything " +
				"about the translation).",
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Type of the IP-binding action\n  * regular - performs One-to-One NAT according to the rule, translates " +
				"the address to to-address;\n  * bypassed - performs the translation, but excludes client from login to the " +
				"HotSpot;\n  * blocked - translation is not performed and packets from a host are dropped.",
			ValidateFunc: validation.StringInSlice([]string{"blocked", "bypassed", "regular"}, false),
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
