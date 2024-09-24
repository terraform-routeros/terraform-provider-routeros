package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*6",
  "action": "deny",
  "disabled": "false",
  "dst-host": "1.2.3.4",
  "dst-port": "!123",
  "dynamic": "false",
  "hits": "0",
  "method": "GET",
  "path": "/sss",
  "server": "server1",
  "src-address": "4.3.2.1"
}
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot/Walled_Garden
func ResourceIpHotspotWalledGarden() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/walled-garden"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("hits", "dst_address"),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Action to perform, when packet matches the rule `allow` - allow access to the web-page without " +
				"authorization, `deny` - the authorization is required to access the web-page.",
			ValidateFunc:     validation.StringInSlice([]string{"allow", "deny"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Domain name of the destination web-server.",
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "TCP port number, client sends request to.",
		},
		KeyDynamic: PropDynamicRo,
		"method": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HTTP method of the request.",
		},
		"path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The path of the request, path comes after `http://dst_host/`.",
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
