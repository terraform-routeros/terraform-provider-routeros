package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
[
        {
        ".id": "*2",
        "name": "netboot",
        "options": "tftpserver-66,unifi,mtu-jumbo"
    }
]
*/

// ResourceDhcpServerOption https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Server
func ResourceDhcpServerOptionSet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/option/sets"),
		MetaId:           PropId(Id),
		KeyComment:       PropCommentRw,
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the DHCP option",
		},
		"options": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The comma sepparated list of options",
		},
	}
	return &schema.Resource{
		Description: "Creates a DHCP lease on the mikrotik device.",

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
