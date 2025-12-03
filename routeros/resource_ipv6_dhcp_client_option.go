package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceDhcpClient https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPClient
func ResourceIPv6DhcpClientOption() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-client/option"),
		MetaId:           PropId(Id),
		KeyName:          PropName("The name that will be used in dhcp-client."),
		"code": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The dhcp-client option code.",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The dhcp-client option",
		},
		"raw_value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "raw_value is computed from value.",
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
