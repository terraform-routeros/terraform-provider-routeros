package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
REST JSON
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841835/CRS1xx+and+2xx+series+switches#CRS1xxand2xxseriesswitches-VLAN
func ResourceInterfaceEthernetSwitchCrsVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/vlan"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"flood": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables or disables forced VLAN flooding per VLAN. If the feature is enabled, the result of " +
				"the destination MAC lookup in the UFDB or MFDB is ignored,and the packet is forced to flood in the VLAN.",
		},
		"ingress_mirror": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable the ingress mirror per VLAN to support the VLAN-based mirror function.",
		},
		"learn": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables or disables source MAC learning for VLAN.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ports": {
			Type:     schema.TypeString,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Member ports of the VLAN.",
		},
		"qos_group": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Defined QoS group from QoS group menu.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"svl": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "FDB lookup mode for lookup in UFDB and MFDB.\n    -  Shared VLAN Learning (svl) - learning/lookup is " +
				"based on MAC addresses - not on VLAN IDs.\n    -  Independent VLAN Learning (ivl) - learning/lookup is based " +
				"on both MAC addresses and VLAN IDs.",
		},
		KeyVlanId: PropVlanIdRw("VLAN ID of the VLAN member entry.", true),
	}

	return &schema.Resource{
		Description:   "Resource for managing CRS (Cloud Router Switch) series device properties.",
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
