package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "disabled": "false",
  "dscp": "46",
  "hw-offloaded": "true",
  "inactive": "false",
  "name": "1-Real-Time (EF/VoIP/Video)",
  "pcp": "0",
  "traffic-class": "7"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CRS3xx,+CRS5xx,+CCR2116,+CCR2216+QoS
// ResourceInterfaceEthernetSwitchQosProfile maps a DSCP/PCP value to an internal
// switch-chip traffic class. Profiles are the building block referenced by the
// `qos map ip` entries and by the per-port `qos port` configuration.
func ResourceInterfaceEthernetSwitchQosProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/profile"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("hw_offloaded", "inactive"),

		KeyName: PropName("Profile name, referenced from the `qos map ip` and `qos port` menus."),
		"dscp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Differentiated Services Code Point (DSCP) value assigned to this profile.",
			ValidateFunc:     validation.IntBetween(0, 63),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Priority Code Point (802.1p) value assigned to this profile.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"traffic_class": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Internal switch-chip traffic class (egress queue) that traffic matching this profile is mapped to.",
			ValidateFunc:     validation.IntBetween(0, 7),
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
