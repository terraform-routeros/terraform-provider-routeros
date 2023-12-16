package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "disabled": "false",
    "name": "steering1",
    "neighbor-group": "something",
    "rrm": "true",
    "wnm": "true"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Steeringproperties
func ResourceWifiSteering() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/steering"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyName: PropName("Name of the steering profile."),
		"neighbor_group": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Neighbor group of potential roaming candidates.",
		},
		"rrm": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable sending 802.11k neighbor reports.",
		},
		"wnm": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable sending 802.11v BSS transition management requests.",
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
