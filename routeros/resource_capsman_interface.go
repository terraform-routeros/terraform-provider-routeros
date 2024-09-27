package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "arp-timeout": "auto",
  "bound": "true",
  "configuration": "test",
  "current-authorized-clients": "0",
  "current-basic-rate-set": "OFDM:6",
  "current-channel": "2462/20/gn(30dBm)",
  "current-rate-set": "OFDM:6-54 BW:1x SGI:1x HT:0-15",
  "current-registered-clients": "0",
  "current-state": "running-ap",
  "disabled": "false",
  "inactive": "false",
  "l2mtu": "1600",
  "mac-address": "00:00:00:00:00:00",
  "master": "true",
  "master-interface": "none",
  "name": "cap1",
  "radio-mac": "00:00:00:00:00:00",
  "radio-name": "000000000000",
  "running": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/interface"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("current_authorized_clients", "current_basic_rate_set", "current_channel",
			"current_rate_set", "current_registered_clients", "current_state"),
		MetaTransformSet: PropTransformSet("channel.config: channel", "datapath.config: datapath",
			"rates.config: rates", "security.config: security"),

		KeyArpTimeout: PropArpTimeoutRw,
		"bound": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface is currently available for the CAPsMAN.",
		},
		"channel": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Channel inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"configuration": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Configuration inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"datapath": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Datapath inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"inactive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface is currently inactive.",
		},
		KeyL2Mtu: PropL2MtuRo,
		"mac_address": {
			Type:             schema.TypeString,
			Description:      "MAC address (BSSID) to use for the interface.",
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"master": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface is not a virtual one.",
		},
		"master_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "The corresponding master interface of the virtual one.",
		},
		KeyName: PropName("Name of the interface."),
		"radio_mac": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The MAC address of the associated radio.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radio_name": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the associated radio.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rates": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Rates inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"running": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface has established a link to another device.",
		},
		"security": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Security inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
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
