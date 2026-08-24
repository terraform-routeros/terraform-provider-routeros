package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*0",
  "disabled": "false",
  "dscp": "46",
  "hw-offloaded": "true",
  "inactive": "false",
  "map": "default",
  "profile": "1-Real-Time (EF/VoIP/Video)"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CRS3xx,+CRS5xx,+CCR2116,+CCR2216+QoS
// ResourceInterfaceEthernetSwitchQosMapIp maps an ingress DSCP value to a QoS
// profile within a named map (usually `default`). This is how L3 (IP) traffic
// classification is programmed into the switch chip.
func ResourceInterfaceEthernetSwitchQosMapIp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/map/ip"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("hw_offloaded", "inactive"),

		"dscp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Ingress Differentiated Services Code Point (DSCP) value to match.",
			ValidateFunc:     validation.IntBetween(0, 63),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"profile": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the `qos profile` this DSCP value is mapped to.",
		},
		"map": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "default",
			Description:      "Name of the QoS map set this entry belongs to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
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
