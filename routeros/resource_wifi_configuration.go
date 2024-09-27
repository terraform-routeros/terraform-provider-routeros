package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*3",
    "aaa": "aaa1",
    "antenna-gain": "10",
    "beacon-interval": "1s",
    "chains": "0,1,2,3",
    "channel": "channel1",
    "country": "Netherlands",
    "datapath": "datapath1",
    "disabled": "false",
    "dtim-period": "1",
    "hide-ssid": "true",
    "interworking": "interworking1",
    "manager": "capsman",
    "mode": "ap",
    "multicast-enhance": "disabled",
    "name": "cfg1",
    "qos-classifier": "priority",
    "security": "security1",
    "security.connect-priority": "0",
    "ssid": "test",
    "steering": "steering1",
    "tx-chains": "4,5,6,7",
    "tx-power": "10"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-Configurationproperties
func ResourceWifiConfiguration() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/configuration"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("aaa.config: aaa", "channel.config: channel", "datapath.config: datapath",
			"interworking.config: interworking", "security.config: security", "steering.config: steering"),

		"aaa": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "AAA inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"antenna_gain": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "An option overrides the default antenna gain.",
			ValidateFunc: validation.IntBetween(0, 30),
		},
		"beacon_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time interval between beacon frames.",
			DiffSuppressFunc: TimeEquall,
		},
		"chains": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
			},
			Description: "Radio chains to use for receiving signals.",
		},
		"channel": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Channel inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"country": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An option determines which regulatory domain restrictions are applied to an interface.",
		},
		"datapath": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Datapath inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"dtim_period": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "A period at which to transmit multicast traffic, when there are client devices in power save mode connected to the AP.",
			ValidateFunc: validation.IntBetween(1, 255),
		},
		"hide_ssid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "This property has effect only in AP mode. Setting it to yes can remove this network from " +
				"the list of wireless networks that are shown by some client software. Changing this setting does not " +
				"improve the security of the wireless network, because SSID is included in other frames sent by the AP.",
		},
		"interworking": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Interworking inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"manager": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to specify the remote CAP mode.",
			ValidateFunc: validation.StringInSlice([]string{"capsman", "capsman-or-local", "local"}, false),
		},
		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to specify the access point operational mode.",
			ValidateFunc: validation.StringInSlice([]string{"ap", "station", "station-bridge"}, false),
		},
		"multicast_enhance": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to enable converting every multicast-address IP or IPv6 packet into multiple unicast-addresses frames for each connected station.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
		},
		KeyName: PropName("Name of the configuration."),
		"qos_classifier": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to specify the QoS classifier.",
			ValidateFunc: validation.StringInSlice([]string{"dscp-high-3-bits", "priority"}, false),
		},
		"security": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Security inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SSID (service set identifier) is a name broadcast in the beacons that identifies wireless network.",
		},
		"steering": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Steering inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_chains": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
			},
			Description: "Radio chains to use for transmitting signals.",
		},
		"tx_power": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "A limit on the transmit power (in dBm) of the interface.",
			ValidateFunc: validation.IntBetween(0, 40),
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
