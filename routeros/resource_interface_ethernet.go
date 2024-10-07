package routeros

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"

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

// ResourceInterfaceEthernet is the schema for ethernet interfaces
// https://help.mikrotik.com/docs/display/ROS/Ethernet#Ethernet-Properties
func ResourceInterfaceEthernet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("factory_name", "driver_rx_byte", "driver_rx_packet", "driver_tx_byte", "driver_tx_packet",
			"rx_64", "rx_65_127", "rx_128_255", "rx_256_511", "rx_512_1023", "rx_1024_1518", "rx_1519_max",
			"tx_64", "tx_65_127", "tx_128_255", "tx_256_511", "tx_512_1023", "tx_1024_1518", "tx_1519_max",
			"tx_rx_64", "tx_rx_65_127", "tx_rx_128_255", "tx_rx_256_511", "tx_rx_512_1023", "tx_rx_1024_1518", "tx_rx_1024_max", "tx_rx_1519_max",
			"rx_broadcast", "rx_bytes", "rx_control", "rx_drop", "rx_fcs_error", "rx_fragment", "rx_jabber", "rx_multicast", "rx_packet", "rx_pause", "rx_too_short", "rx_too_long",
			"tx_broadcast", "tx_bytes", "tx_control", "tx_drop", "tx_fcs_error", "tx_fragment", "tx_jabber", "tx_multicast", "tx_packet", "tx_pause", "tx_too_short", "tx_too_long",
			"rx_align_error", "rx_carrier_error", "rx_code_error", "rx_error_events", "rx_length_error", "rx_overflow", "rx_unicast", "rx_unknown_op",
			"tx_collision", "tx_excessive_collision", "tx_late_collision", "tx_multiple_collision", "tx_single_collision", "tx_total_collision",
			"tx_deferred", "tx_excessive_deferred", "tx_unicast", "tx_underrun", "rx_tcp_checksum_error", "rx_udp_checksum_error", "rx_ip_header_checksum_error",
			"tx_carrier_sense_error",
		),

		"advertise": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `
				Advertised speed and duplex modes for Ethernet interfaces over twisted pair, 
				only applies when auto-negotiation is enabled. Advertising higher speeds than 
				the actual interface supported speed will have no effect, multiple options are allowed.`,
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^[0-9\.]+(M|G)-(full|half|base\w+(-\w+)?)$`),
				"Since RouterOS v7.12 the values of this property have changed. Please check the documentation.",
			),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"auto_negotiation": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: `When enabled, the interface "advertises" its maximum capabilities to achieve the best connection possible.
					Note1: Auto-negotiation should not be disabled on one end only, otherwise Ethernet Interfaces may not work properly.
					Note2: Gigabit Ethernet and NBASE-T Ethernet links cannot work with auto-negotiation disabled.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
		KeyComment:     PropCommentRw,
		KeyDefaultName: PropDefaultNameRo("The default name for an interface."),
		KeyDisabled:    PropDisabledRw,
		"disable_running_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: `Disable running check. If this value is set to 'no', the router automatically detects whether the NIC is connected with a device in the network or not.
			Default value is 'yes' because older NICs do not support it. (only applicable to x86)`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"factory_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The factory name of the identifier, serves as resource identifier. Determines which interface will be updated.",
		},
		"fec_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Changes Forward Error Correction (FEC) mode for SFP28, QSFP+ and QSFP28 interfaces. " +
				"Same mode should be used on both link ends, otherwise FEC mismatch could result in non-working link " +
				"or even false link-ups. ",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "fec74", "fec91", "off"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"full_duplex": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      `Defines whether the transmission of data appears in two directions simultaneously, only applies when auto-negotiation is disabled.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyL2Mtu:                   PropL2MtuRw,
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		KeyLoopProtectStatus:       PropLoopProtectStatusRo,
		"mac_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Media Access Control number of an interface.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mdix_enable": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      `Whether the MDI/X auto cross over cable correction feature is enabled for the port (Hardware specific, e.g. ether1 on RB500 can be set to yes/no. Fixed to 'yes' on other hardware.)`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Layer3 Maximum transmission unit",
			ValidateFunc:     validation.IntBetween(0, 65536),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the ethernet interface."),
		"orig_mac_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Original Media Access Control number of an interface. (read only)",
		},
		"poe_lldp_enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option that enables LLDP for managing devices.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"poe_out": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)",
			ValidateFunc:     validation.StringInSlice([]string{"auto-on", "forced-on", "off"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"poe_priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)",
			ValidateFunc:     validation.IntBetween(0, 99),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"poe_voltage": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option that allows us to manually control the voltage outputs on the PoE port.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "high", "low"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"power_cycle_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An options that disables PoE-Out power for 5s between the specified intervals.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"power_cycle_ping_enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option that enables ping watchdog of power cycles on the port if a host does not respond to ICMP or MAC-Telnet packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"power_cycle_ping_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An address to monitor.",
			ValidateFunc: validation.IsIPAddress,
			RequiredWith: []string{"power_cycle_ping_enabled"},
		},
		"power_cycle_ping_timeout": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If the host does not respond over the specified period, the PoE-Out port is switched off for 5s.",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return AlwaysPresentNotUserProvided(k, old, new, d) || TimeEquall(k, old, new, d)
			},
		},
		"running": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether interface is running. Note that some interface does not have running check and they are always reported as \"running\"",
		},
		"rx_flow_control": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `When set to on, the port will process received pause frames and suspend transmission if required.
					auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.`,
			ValidateFunc:     validation.StringInSlice([]string{"on", "off", "auto"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sfp_rate_select": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Allows to control rate select pin for SFP ports. Values: high | low`,
			ValidateFunc:     validation.StringInSlice([]string{"high", "low"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sfp_shutdown_temperature": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The temperature in Celsius at which the interface will be temporarily turned off due to too high detected SFP module temperature (introduced v6.48)." +
				"The default value for SFP/SFP+/SFP28 interfaces is 95, and for QSFP+/QSFP28 interfaces 80 (introduced v7.6).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sfp_ignore_rx_los": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "An option to ignore RX LOS (Loss of Signal) status of the SFP module.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"slave": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether interface is configured as a slave of another interface (for example Bonding)",
		},
		"speed": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets interface data transmission speed which takes effect only when ```auto_negotiation``` is disabled.",
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^[0-9\.]+(M|G)(?:(bps)|(-base\w+)(-\w+)?)$`),
				"Since RouterOS v7.12 the values of this property have changed. Please check the documentation.",
			),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"switch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID to which switch chip interface belongs to.",
		},
		"tx_flow_control": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `When set to on, the port will generate pause frames to the upstream device to temporarily stop the packet transmission. 
					Pause frames are only generated when some routers output interface is congested and packets cannot be transmitted anymore. 
					Auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.`,
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
	var ethernetInterface MikrotikItem
	var err error

	// The item has already an ID, we don't need to perform the lookup
	if val := d.Id(); val != "" {
		tflog.Debug(ctx, "fetching ethernet interface by id", map[string]interface{}{"id": val})
		metadata := GetMetadata(s)
		items, err := ReadItems(&ItemId{metadata.IdType, val}, metadata.Path, m.(Client))

		if err != nil {
			return diag.FromErr(fmt.Errorf("reading interface by id: %w", err))
		}
		if len(*items) > 1 {
			return diag.FromErr(fmt.Errorf("more than 1 interface returned when fetching by id %v", val))
		}
		if len(*items) == 0 {
			return diag.FromErr(fmt.Errorf("unable to find interface when fetching by id: %v", val))
		}
		ethernetInterface = (*items)[0]
	} else {
		// As We don't know the ID, we have to look it up by "default"/"factory" name
		ethernetInterface, err = findInterfaceByDefaultName(s, d, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	s = updateSchemaWithRouterCapabilities(s, ethernetInterface)
	return DefaultRead(s)(ctx, d, m)
}

// updateEthernetInterface searches for the interface and disables fields not supported by the router instance
func updateEthernetInterface(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if val := d.Id(); val == "" {
		ethernetInterface, err := findInterfaceByDefaultName(s, d, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(ethernetInterface.GetID(Id))
	}

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
