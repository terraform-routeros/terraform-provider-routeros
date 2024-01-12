package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*2",
  "disabled": "false",
  "invalid": "false",
  "ports": "ether1",
  "switch": "switch1",
  "vlan-id": "0"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features#SwitchChipFeatures-VLANTable
func ResourceInterfaceEthernetSwitchVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/interface/ethernet/switch/vlan"),
		MetaId:             PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"independent_learning": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to use shared-VLAN-learning (SVL) or independent-VLAN-learning (IVL).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInvalid: PropInvalidRo,
		"ports": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Interface member list for the respective VLAN.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: AlwaysPresentNotUserProvided,
			},
		},
		"switch": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the switch for which the respective VLAN entry is intended for.",
		},
		"vlan_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The VLAN ID for certain switch port configurations.",
			ValidateFunc:     validation.IntBetween(0, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
