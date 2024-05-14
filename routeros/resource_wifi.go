package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".about": "mode: AP, SSID: wlan, channel: 2462/n",
    ".id": "*1",
    "arp": "enabled",
    "arp-timeout": "auto",
    "bound": "true",
    "configuration": "cfg1",
    "configuration.manager": "capsman",
    "configuration.mode": "ap",
    "default-name": "wifi1",
    "disabled": "false",
    "inactive": "false",
    "l2mtu": "1560",
    "mac-address": "00:00:00:00:00:00",
    "master": "true",
    "name": "wifi1",
    "radio-mac": "00:00:00:00:00:00",
    "running": "true",
    "security.connect-priority": "0"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Miscellaneousproperties
// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Read-onlyproperties
func ResourceWifi() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("aaa: aaa.config", "channel: channel.config", "configuration: configuration.config",
			"datapath: datapath.config", "interworking: interworking.config", "security: security.config", "steering: steering.config"),

		"aaa": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "AAA inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"bound": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface is currently available for the WiFi manager.",
		},
		"channel": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Channel inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"configuration": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Configuration inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"datapath": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Datapath inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"default_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The interface's default name.",
		},
		KeyDisabled: PropDisabledRw,
		"disable_running_check": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to set the running property to true if it is not disabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"inactive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the interface is currently inactive.",
		},
		"interworking": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Interworking inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyL2Mtu: PropL2MtuRw,
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
			Description: "The corresponding master interface of the virtual one.",
		},
		"mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Layer3 maximum transmission unit",
			ValidateFunc:     validation.IntBetween(32, 2290),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the interface."),
		"radio_mac": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The MAC address of the associated radio.",
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
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"steering": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Steering inline settings.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
