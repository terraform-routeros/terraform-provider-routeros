package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceDhcpClient https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPClient
func ResourceDhcpClientOption() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-client/option"),
		MetaId:           PropId(Id),
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name that will be used in dhcp-client.",
		},
		"code": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The dhcp-client option code.",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:     true,
			Description: "The dhcp-client option",
		},
		"raw_value":
		{
			Type:        schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description: "raw_value is computed from value.",
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
