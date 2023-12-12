package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "bridge": "lan",
    "bridge-cost": "1",
    "bridge-horizon": "none",
    "client-isolation": "true",
    "disabled": "false",
    "interface-list": "LAN",
    "name": "datapath1",
    "vlan-id": "1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Datapathproperties
func ResourceWifiDatapath() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/datapath"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge interface to add the interface as a bridge port.",
		},
		"bridge_cost": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Spanning tree protocol cost of the bridge port.",
		},
		"bridge_horizon": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge horizon to use when adding as a bridge port.",
		},
		"client_isolation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to toggle communication between clients connected to the same AP.",
		},
		KeyComment: PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List to which add the interface as a member.",
		},
		KeyName: PropName("Name of the datapath."),
		"vlan_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Default VLAN ID to assign to client devices connecting to this interface.",
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
