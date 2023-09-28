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

// https://help.mikrotik.com/docs/display/ROS/Ethernet#Ethernet-Properties
func ResourceInterfaceEthernet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields(`"factory_name"`),
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
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "default",
			Description:  `Changes the cable length setting (only applicable to NS DP83815/6 cards)`,
			ValidateFunc: validation.StringInSlice([]string{"default", "short", "standard"}, false),
		},
		"combo_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "auto",
			Description: `When auto mode is selected, the port that was first connected will establish the link. In case this link fails, the other port will try to establish a new link. If both ports are connected at the same time (e.g. after reboot), 
				the priority will be the SFP/SFP+ port. When sfp mode is selected, the interface will only work through SFP/SFP+ cage.
				When copper mode is selected, the interface will only work through RJ45 Ethernet port.`,
			ValidateFunc: validation.StringInSlice([]string{"auto", "copper", "sfp"}, false),
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
		"driver_rx_byte": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received bytes on device CPU`,
		},
		"driver_rx_packet": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received packets on device CPU`,
		},
		"driver_tx_byte": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of transmitted packets by device CPU`,
		},
		"driver_tx_packet": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of transmitted packets by device CPU`,
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
		KeyL2Mtu: PropL2MtuRo,
		"loop_protect": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "default",
			ValidateFunc: validation.StringInSlice([]string{"default", "on", "off"}, false),
		},
		"loop_protect_disable_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "5m",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"loop_protect_send_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "5s",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"loop_protect_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"mac_address": {
			Type:        schema.TypeString,
			Description: `Media Access Control number of an interface.`,
			Optional:    true,
			Computed:    true,
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
			Type:         schema.TypeString,
			Description:  "PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)",
			Default:      "off",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"auto-on", "forced-on", "off"}, false),
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
		"rx_broadcast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received broadcast frames.",
		},
		"rx_bytes": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received bytes.",
		},
		"rx_error_events": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received frames with active error event`,
		},
		"rx_fcs_error": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received frames with incorrect checksum",
		},
		"rx_fragment": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received fragmented frames (not related to IP fragmentation)`,
		},
		"rx_jabber": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received jabbed packets - a packet that is transmitted longer than the maximum packet length",
		},
		"rx_multicast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received multicast frames.",
		},
		"rx_packet": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received packets.",
		},
		"rx_pause": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of received pause frames",
		},
		"rx_flow_control": {
			Type: schema.TypeString,
			Description: `When set to on, the port will process received pause frames and suspend transmission if required.
					auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.`,
			Default:      "off",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"on", "off", "auto"}, false),
		},
		"rx_overflow": {
			Type:        schema.TypeInt,
			Description: `Total count of received overflowed frames, can be caused when device resources are insufficient to receive a certain frame`,
			Computed:    true,
		},
		"rx_too_long": {
			Type:        schema.TypeInt,
			Description: `Total count of received frames that were larger than the maximum supported frame size by the network device, see the max-l2mtu property`,
			Computed:    true,
		},
		"rx_too_short": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received frame shorter than the minimum 64 bytes`,
		},
		"rx_unicast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: `Total count of received unicast frames`,
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
			Default:      "off",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"on", "off", "auto"}, false),
		},
		"tx_broadcast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted broadcast frames.",
		},
		"tx_bytes": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted bytes.",
		},
		"tx_multicast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted multicast frames.",
		},
		"tx_collision": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted frames that made collisions",
		},
		"tx_drop": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted frames that were dropped due to already full output queue",
		},

		"tx_late_collision": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted frames that made collision after being already halfway transmitted",
		},
		"tx_packet": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted packets.",
		},
		"tx_pause": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted pause frames.",
		},
		"tx_rx_64": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 64 byte frames",
		},
		"tx_rx_65_127": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 64 to 127 byte frames",
		},
		"tx_rx_128_255": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 128 to 255 byte frames",
		},
		"tx_rx_256_511": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 256 to 511 byte frames",
		},
		"tx_rx_512_1023": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 512 to 1024 byte frames",
		},
		"tx_rx_1024_max": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted and received 1024 or above byte frames",
		},
		"tx_underrun": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted underrun packets",
		},
		"tx_unicast": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total count of transmitted unicast frames.",
		},
	}

	return &schema.Resource{
		CreateContext: UpdateOnlyDeviceCreate(resSchema),
		ReadContext:   UpdateOnlyDeviceRead(resSchema),
		UpdateContext: UpdateOnlyDeviceUpdate(resSchema),
		DeleteContext: NoOpDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}

func UpdateOnlyDeviceCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return UpdateEthernetInterface(ctx, s, d, m)
	}
}

func UpdateOnlyDeviceUpdate(s map[string]*schema.Schema) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return UpdateEthernetInterface(ctx, s, d, m)
	}
}

func UpdateOnlyDeviceRead(s map[string]*schema.Schema) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return DefaultRead(s)(ctx, d, m)
	}
}

func UpdateEthernetInterface(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ethernetInterface, err := findInterfaceByDefaultName(s, d, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// Router won't accept poe-out parameter if the interface does not support it.
	poeDesiredState := d.Get(poeOutField)
	_, supportsPoE := ethernetInterface[SnakeToKebab(poeOutField)]
	switch {
	// if the user has specified it, but it's not supported, let's error out
	case poeDesiredState != "off" && !supportsPoE:
		return diag.FromErr(errors.New("can't configure PoE, router does not supports it"))
	// if the router does not support PoE, avoid sending the parameter as it returns an error.
	case !supportsPoE:
		skipFieldInSchema(s, poeOutField)
	}

	if _, supportsCableSettings := ethernetInterface["cable-settings"]; supportsCableSettings {
		skipFieldInSchema(s, cableSettingsField)
	}

	if _, supportsRunningCheck := ethernetInterface["disable-running-check"]; supportsRunningCheck {
		skipFieldInSchema(s, runningCheckField)
	}

	// Dynamic schema, counters for tx_queue packets, changes from router to router, read only countes.
	for key, _ := range ethernetInterface {
		if strings.HasPrefix(key, "tx_queue") {
			fmt.Printf("adding dynamic schema: %s", key)
			s[key] = &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total count of interface for numbered queue",
			}
		}
	}

	d.SetId(ethernetInterface.GetID(Id))
	if updateDiag := ResourceUpdate(ctx, s, d, m); updateDiag.HasError() {
		return updateDiag
	}

	return ResourceRead(ctx, s, d, m)

}

func NoOpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
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

func skipFieldInSchema(s map[string]*schema.Schema, field string) {
	s[MetaSkipFields].Default = fmt.Sprintf("%s,\"%s\"", s[MetaSkipFields].Default, field)
}
