package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
".id":"*8",
"arp":"enabled",
"arp-timeout":"auto",
"beamforming-event":"3042",
"disabled":"false",
"mac-address":"48:8F:5A:F3:34:E0",
"mtu":"1500",
"name":"ptp01",
"parent":"wlan60-1",
"put-in-bridge":"parent",
"remote-address":"48:8F:5A:47:81:9E",
"running":"true"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/39059501/W60G
func ResourceInterfaceW60gStation() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/w60g/station"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("beamforming_event"),

		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		KeyDisabled:   PropDisabledRw,
		KeyMacAddress: PropMacAddressRw("MAC address of the station interface.", false),
		KeyMtu:        PropMtuRw(),
		KeyName:       PropName("Name of the interface."),
		"parent": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Parent interface name.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"put_in_bridge": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Add station device interface to specific bridge.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "MAC address of bridge interface, station is connecting to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
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
