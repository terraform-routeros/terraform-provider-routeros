package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*4",
  "action": "reject",
  "disabled": "false",
  "dst-address": "!0.0.0.0",
  "dst-address-list": "bbb",
  "dst-port": "0-65535",
  "invalid": "false",
  "protocol": "tcp",
  "server": "server1",
  "src-address": "0.0.0.0",
  "src-address-list": "aaa"
}
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot/Walled_Garden#IP_Walled_Garden
func ResourceIpHotspotWalledGardenIp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/walled-garden/ip"),
		MetaId:           PropId(Id),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Action to perform, when packet matches the rule allow - allow access to the web-page without " +
				"authorization deny - the authorization is required to access the web-page reject - the authorization " +
				"is required to access the resource, ICMP reject message will be sent to client, when packet will match " +
				"the rule.",
			ValidateFunc:     validation.StringInSlice([]string{"allow", "deny", "reject"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "Destination IP address, IP address of the WEB-server. Ignored if dst-host is already specified.",
			ConflictsWith: []string{"dst_host"},
		},
		"dst_address_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination IP address list. Ignored if dst-host is already specified.",
		},
		"dst_host": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Domain name of the destination web-server. When this parameter is specified dynamic entry " +
				"is added to Walled Garden.",
			ConflictsWith: []string{"dst_address"},
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "TCP port number, client sends request to.",
		},
		KeyInvalid: PropInvalidRo,
		"protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP protocol.",
		},
		"server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the HotSpot server, rule is applied to.",
		},
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Source address of the user, usually IP address of the HotSpot client.",
		},
		"src_address_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Source IP address list.",
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
