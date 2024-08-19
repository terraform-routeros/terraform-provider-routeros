package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
[
    {
        ".id": "*4",
        "code": "66",
        "name": "optionname",
        "raw-value": "00002311",
        "value": "0x00002311"
    }
]
*/

// ResourceDhcpServerOption https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPServer
func ResourceDhcpServerOption() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/option"),
		MetaId:           PropId(Id),
		KeyComment:       PropCommentRw,
		"code": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "The number of the DHCP option",
			ValidateFunc: validation.IntBetween(1, 254),
		},
		"force": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Force the DHCP option from the server-side even if the DHCP-client does not request such parameter.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the DHCP option",
		},
		"raw_value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The computed value of the option as an hex value",
		},
		"value": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The value with formatting using Mikrotik settings https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Server#DHCP_Options",
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
