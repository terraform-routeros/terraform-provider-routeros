package routeros

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "advertise": "10M-half,10M-full,100M-half,100M-full,1000M-full",
    "arp": "enabled",
    "arp-timeout": "auto",
    "auto-negotiation": "true",
    "cable-settings": "default",
    "default-name": "ether1",
    "disable-running-check": "true",
    "disabled": "false",
    "loop-protect": "default",
    "loop-protect-disable-time": "5m",
    "loop-protect-send-interval": "5s",
    "loop-protect-status": "off",
    "mac-address": "54:05:AB:1E:BE:71",
    "mtu": "1500",
    "name": "ether1",
    "orig-mac-address": "54:05:AB:1E:BE:71",
    "running": "true",
    "rx-broadcast": "250",
    "rx-bytes": "222253",
    "rx-flow-control": "off",
    "rx-multicast": "10",
    "rx-packet": "1889",
    "tx-broadcast": "113",
    "tx-bytes": "693931",
    "tx-flow-control": "off",
    "tx-multicast": "270",
    "tx-packet": "2222"
}
*/

const poeOutField = "poe_out"
const cableSettingsField = "cable_settings"
const runningCheckField = "disable_running_check"

// ResourceInterfaceEthernet is the schema for ethernet interfaces
// https://help.mikrotik.com/docs/display/ROS/Ethernet#Ethernet-Properties
func ResourceInterfaceEthernet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields(`"factory_name","driver_rx_byte","driver_rx_packet","driver_tx_byte","driver_tx_packet",` +
			`"rx_64","rx_65_127","rx_128_255","rx_256_511","rx_512_1023","rx_1024_1518","rx_1519_max",` +
			`"tx_64","tx_65_127","tx_128_255","tx_256_511","tx_512_1023","tx_1024_1518","tx_1519_max",` +
			`"tx_rx_64","tx_rx_65_127","tx_rx_128_255","tx_rx_256_511","tx_rx_512_1023","tx_rx_1024_1518","tx_rx_1519_max",` +
			`"rx_broadcast","rx_bytes","rx_control","rx_drop","rx_fcs_error","rx_fragment","rx_jabber","rx_multicast","rx_packet","rx_pause","rx_too_short","rx_too_long",` +
			`"tx_broadcast","tx_bytes","tx_control","tx_drop","tx_fcs_error","tx_fragment","tx_jabber","tx_multicast","tx_packet","tx_pause","tx_too_short","tx_too_long",` +
			`"rx_align_error","rx_carrier_error","rx_code_error","rx_length_error","rx_overflow","rx_unknown_op",` +
			`"tx_collision","tx_excessive_collision","tx_late_collision","tx_multiple_collision","tx_single_collision","tx_total_collision",` +
			`"tx_deferred","tx_excessive_deferred","tx_underrun",`),

		"advertise": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
			Description: `
				Advertised speed and duplex modes for Ethernet interfaces over twisted pair, 
				only applies when auto-negotiation is enabled. Advertising higher speeds than 
				the actual interface supported speed will have no effect, multiple options are allowed.`,
			ValidateFunc: validation.StringInSlice([]string{
				"10M-full", "10M-half", "100M-full", "100M-half",
				"1000M-full", "1000M-half", "2500M-full", "5000M-full", "10000M-full"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"auto_negotiation": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: `When enabled, the interface "advertises" its maximum capabilities to achieve the best connection possible.
					Note1: Auto-negotiation should not be disabled on one end only, otherwise Ethernet Interfaces may not work properly.
					Note2: Gigabit Ethernet and NBASE-T Ethernet links cannot work with auto-negotiation disabled.`,
		},
		"bandwidth": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Sets max rx/tx bandwidth in kbps that will be handled by an interface. TX limit is supported on all Atheros switch-chip ports. 
				RX limit is supported only on Atheros8327/QCA8337 switch-chip ports.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cable_settings": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Changes the cable length setting (only applicable to NS DP83815/6 cards)`,
			ValidateFunc:     validation.StringInSlice([]string{"default", "short", "standard"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"combo_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `When auto mode is selected, the port that was first connected will establish the link. In case this link fails, the other port will try to establish a new link. If both ports are connected at the same time (e.g. after reboot), 
				the priority will be the SFP/SFP+ port. When sfp mode is selected, the interface will only work through SFP/SFP+ cage.
				When copper mode is selected, the interface will only work through RJ45 Ethernet port.`,
			ValidateFunc:     validation.StringInSlice([]string{"auto", "copper", "sfp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"default_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The default name for an interface.",
		},
		KeyDisabled: PropDisabledRw,
		"disable_running_check": {
			Type: schema.TypeBool,
			Description: `Disable running check. If this value is set to 'no', the router automatically detects whether the NIC is connected with a device in the network or not.
			Default value is 'yes' because older NICs do not support it. (only applicable to x86)`,
			Default:  true,
			Optional: true,
		},
		"factory_name": {
			Type:        schema.TypeString,
			Optional:    false,
			Required:    true,
			Description: "The factory name of the identifier, serves as resource identifier. Determines which interface will be updated.",
		},
		"full_duplex": {
			Type:        schema.TypeBool,
			Description: `Defines whether the transmission of data appears in two directions simultaneously, only applies when auto-negotiation is disabled.`,
			Default:     true,
			Optional:    true,
		},
		KeyL2Mtu: {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Layer2 Maximum transmission unit. " +
				"[See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).",
		},
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		KeyLoopProtectStatus:       PropLoopProtectStatusRo,
		"mac_address": {
			Type:             schema.TypeString,
			Description:      `Media Access Control number of an interface.`,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mdix_enable": {
			Type:        schema.TypeBool,
			Description: `Whether the MDI/X auto cross over cable correction feature is enabled for the port (Hardware specific, e.g. ether1 on RB500 can be set to yes/no. Fixed to 'yes' on other hardware.)`,
			Optional:    true,
			Default:     true,
		},
		"mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1500,
			Description:  "Layer3 Maximum transmission unit",
			ValidateFunc: validation.IntBetween(0, 65536),
		},
		KeyName: PropName("Name of the ethernet interface."),
		"orig_mac_address": {
			Type:        schema.TypeString,
			Description: "Original Media Access Control number of an interface. (read only)",
			Computed:    true,
		},
		poeOutField: {
			Type:             schema.TypeString,
			Description:      "PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)",
			Default:          "off",
			Optional:         true,
			ValidateFunc:     validation.StringInSlice([]string{"auto-on", "forced-on", "off"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"poe_priority": {
			Type:         schema.TypeInt,
			Description:  "PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)",
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 99),
		},
		"running": {
			Type:        schema.TypeBool,
			Description: "Whether interface is running. Note that some interface does not have running check and they are always reported as \"running\"",
			Computed:    true,
		},
		"rx_flow_control": {
			Type: schema.TypeString,
			Description: `When set to on, the port will process received pause frames and suspend transmission if required.
					auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.`,
			Default:          "off",
			Optional:         true,
			ValidateFunc:     validation.StringInSlice([]string{"on", "off", "auto"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sfp_rate_select": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  `Allows to control rate select pin for SFP ports. Values: high | low`,
			Default:      "high",
			ValidateFunc: validation.StringInSlice([]string{"high", "low"}, false),
		},
		"sfp_shutdown_temperature": {
			Type: schema.TypeInt,
			Description: "The temperature in Celsius at which the interface will be temporarily turned off due to too high detected SFP module temperature (introduced v6.48)." +
				"The default value for SFP/SFP+/SFP28 interfaces is 95, and for QSFP+/QSFP28 interfaces 80 (introduced v7.6).",
			Optional: true,
		},
		"slave": {
			Type:        schema.TypeBool,
			Description: "Whether interface is configured as a slave of another interface (for example Bonding)",
			Computed:    true,
		},
		"speed": {
			Type:         schema.TypeString,
			Description:  "Sets interface data transmission speed which takes effect only when auto-negotiation is disabled.",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"10Mbps", "10Gbps", "100Mbps", "1Gbps"}, false),
		},
		"switch": {
			Type:        schema.TypeString,
			Description: "ID to which switch chip interface belongs to.",
			Computed:    true,
		},
		"tx_flow_control": {
			Type: schema.TypeString,
			Description: `When set to on, the port will generate pause frames to the upstream device to temporarily stop the packet transmission. 
					Pause frames are only generated when some routers output interface is congested and packets cannot be transmitted anymore. 
					Auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.`,
			Default:          "off",
			Optional:         true,
			ValidateFunc:     validation.StringInSlice([]string{"on", "off", "auto"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: updateOnlyDeviceCreate(resSchema),
		ReadContext:   updateOnlyDeviceRead(resSchema),
		UpdateContext: updateOnlyDeviceUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}

func updateOnlyDeviceCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return updateEthernetInterface(ctx, s, d, m)
	}
}

func updateOnlyDeviceUpdate(s map[string]*schema.Schema) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return updateEthernetInterface(ctx, s, d, m)
	}
}

func updateOnlyDeviceRead(s map[string]*schema.Schema) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return readEthernetInterface(ctx, s, d, m)
	}
}

func readEthernetInterface(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ethernetInterface, err := findInterfaceByDefaultName(s, d, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}
	s = updateSchemaWithRouterCapabilities(s, ethernetInterface)
	return DefaultRead(s)(ctx, d, m)
}

// updateEthernetInterface searches for the interface and disables fields not supported by the router instance
func updateEthernetInterface(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ethernetInterface, err := findInterfaceByDefaultName(s, d, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// Router won't accept poe-out parameter if the interface does not support it.
	poeDesiredState := d.Get(poeOutField)
	_, supportsPoE := ethernetInterface[SnakeToKebab(poeOutField)]
	switch {
	// if the user has specified it, but it's not supported, lets error out
	case poeDesiredState != "off" && !supportsPoE:
		return diag.FromErr(errors.New("can't configure PoE, router does not supports it"))
	// if the router does not support PoE, avoid sending the parameter as it returns an error.
	case !supportsPoE:
		s[MetaSkipFields].Default = skipFieldInSchema(s[MetaSkipFields].Default, poeOutField)
	}

	if _, supportsCableSettings := ethernetInterface[SnakeToKebab(cableSettingsField)]; !supportsCableSettings {
		s[MetaSkipFields].Default = skipFieldInSchema(s[MetaSkipFields].Default, cableSettingsField)
	}

	if _, supportsRunningCheck := ethernetInterface[SnakeToKebab(runningCheckField)]; !supportsRunningCheck {
		s[MetaSkipFields].Default = skipFieldInSchema(s[MetaSkipFields].Default, runningCheckField)
	}

	d.SetId(ethernetInterface.GetID(Id))
	if updateDiag := ResourceUpdate(ctx, s, d, m); updateDiag.HasError() {
		return updateDiag
	}

	return readEthernetInterface(ctx, s, d, m)
}

func updateSchemaWithRouterCapabilities(s map[string]*schema.Schema, item MikrotikItem) map[string]*schema.Schema {
	// Dynamic schema, counters for tx_queue${number}_packets, changes from router to router, read only counters.
	// Just drop them as they don't have much sense in the context of a terraform provider
	for key := range item {
		if strings.HasPrefix(key, "tx-queue") {
			s[MetaSkipFields].Default = skipFieldInSchema(s[MetaSkipFields].Default, KebabToSnake(key))
		}
	}
	return s
}

func findInterfaceByDefaultName(s map[string]*schema.Schema, d *schema.ResourceData, c Client) (MikrotikItem, error) {
	metadata := GetMetadata(s)
	filter := buildReadFilter(map[string]interface{}{"default-name": d.Get("factory_name")})
	items, err := ReadItemsFiltered(filter, metadata.Path, c)
	if err != nil {
		return nil, err
	}

	if items == nil || len(*items) == 0 {
		return nil, errors.New("unable to find interface")
	}

	ethernetInterface := (*items)[0]
	return ethernetInterface, nil
}

func skipFieldInSchema(defaults interface{}, field string) string {
	return fmt.Sprintf("%s,\"%s\"", defaults, field)
}
