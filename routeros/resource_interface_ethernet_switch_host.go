package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "copy-to-cpu": "false",
  "drop": "false",
  "dynamic": "false",
  "invalid": "false",
  "mac-address": "00:00:00:00:00:00",
  "mirror": "false",
  "ports": "ether1",
  "redirect-to-cpu": "false",
  "switch": "switch1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features#SwitchChipFeatures-HostTable
func ResourceInterfaceEthernetSwitchHost() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/host"),
		MetaId:           PropId(Id),

		"copy_to_cpu": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send a frame copy to switch CPU port from a frame with matching MAC destination address " +
				"(matching destination or source address for CRS3xx series switches).",
		},
		"drop": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to drop a frame with matching MAC source address received on a certain port (matching " +
				"destination or source address for CRS3xx series switches).",
		},
		KeyDynamic:    PropDynamicRo,
		KeyInvalid:    PropInvalidRo,
		KeyMacAddress: PropMacAddressRw("Host's MAC address.", true),
		"mirror": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send a frame copy to mirror-target port from a frame with matching MAC destination address " +
				"(matching destination or source address for CRS3xx series switches).",
		},
		"ports": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Name of the interface, static MAC address can be mapped to more that one port, including switch CPU port.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: AlwaysPresentNotUserProvided,
			},
		},
		"redirect_to_cpu": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to redirect a frame to switch CPU port from a frame with matching MAC destination address " +
				"(matching destination or source address for CRS3xx series switches).",
		},
		"share_vlan_learned": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the static host MAC address lookup is used with shared-VLAN-learning (SVL) or " +
				"independent-VLAN-learning (IVL). The SVL mode is used for those VLAN entries that do not support IVL or IVL is " +
				"disabled (independent-learning=no).",
		},
		"switch": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the switch to which the MAC address is going to be assigned to.",
		},
		"vlan_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN ID for the statically added MAC address entry.",
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
