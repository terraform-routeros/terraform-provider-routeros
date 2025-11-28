package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*2",
  "comment": "",
  "disabled": "false",
  "dynamic": "false",
  "tagged-ports": "ether1,ether2,ether3,ether4,ether5,ether6",
  "vlan-id": "100"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841835/CRS1xx+and+2xx+series+switches#CRS1xxand2xxseriesswitches-EgressVLANTag
func ResourceInterfaceEthernetSwitchCrsEgressVlanTag() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/egress-vlan-tag"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"tagged_ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Ports that are tagged in egress.",
		},
		"vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "VLAN ID which is tagged in egress.",
			ValidateFunc: validation.IntBetween(1, 4095),
		},
	}

	return &schema.Resource{
		Description: "Resource for managing CRS (Cloud Router Switch) series device properties.",

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
