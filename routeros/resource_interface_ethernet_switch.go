package routeros

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
	".id":"*0",
	"cpu-flow-control":"true",
	"driver-rx-byte":"47500186966",
	"driver-rx-packet":"107352097",
	"driver-tx-byte":"202118920815",
	"driver-tx-packet":"175081916",
	"invalid":"false",
	"mirror-source":"none",
	"mirror-target":"none",
	"name":"switch1",
	"rx-align-error":"0",
	"rx-broadcast":"1049998",
	"rx-bytes":"55163404001",
	"rx-carrier-error":"0",
	"rx-code-error":"0",
	"rx-control":"0",
	"rx-drop":"0",
	"rx-fcs-error":"0",
	"rx-fragment":"0",
	"rx-jabber":"0",
	"rx-length-error":"0",
	"rx-multicast":"5101395",
	"rx-packet":"145932800",
	"rx-pause":"0",
	"rx-too-long":"75",
	"rx-too-short":"0",
	"rx-unknown-op":"0",
	"tx-broadcast":"276405",
	"tx-bytes":"202693687261",
	"tx-control":"0",
	"tx-deferred":"0",
	"tx-drop":"0",
	"tx-excessive-collision":"0",
	"tx-excessive-deferred":"0",
	"tx-fcs-error":"0",
	"tx-fragment":"0",
	"tx-jabber":"0",
	"tx-late-collision":"0",
	"tx-multicast":"737076",
	"tx-multiple-collision":"0",
	"tx-packet":"175076294",
	"tx-pause":"0",
	"tx-rx-1024-1518":"169188042",
	"tx-rx-128-255":"45589148",
	"tx-rx-1519-max":"0",
	"tx-rx-256-511":"4444935",
	"tx-rx-512-1023":"3890955",
	"tx-rx-64":"5245202",
	"tx-rx-65-127":"92433137",
	"tx-single-collision":"0",
	"tx-too-long":"218648",
	"tx-too-short":"0",
	"tx-total-collision":"0",
	"type":"QCA-8337"
}

[=Oo=@mikrotik]/interface/ethernet/switch/print detail
Flags: I - invalid
 0   name="switch1" type=QCA-8337 mirror-source=none mirror-target=none mirror-egress-target=none cpu-flow-control=yes l3-hw-offloading=no

*/

// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features#SwitchChipFeatures-CPUFlowControl
// https://help.mikrotik.com/docs/display/ROS/L3+Hardware+Offloading
func ResourceInterfaceEthernetSwitch() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("switch_id", "driver_rx_byte", "driver_rx_packet", "driver_tx_byte", "driver_tx_packet",
			"rx_align_error", "rx_broadcast", "rx_bytes", "rx_carrier_error", "rx_code_error", "rx_control", "rx_drop",
			"rx_fcs_error", "rx_fragment", "rx_jabber", "rx_length_error", "rx_multicast", "rx_packet", "rx_pause",
			"rx_too_long", "rx_too_short", "rx_unknown_op", "tx_broadcast", "tx_bytes", "tx_control", "tx_deferred",
			"tx_drop", "tx_excessive_collision", "tx_excessive_deferred", "tx_fcs_error", "tx_fragment", "tx_jabber",
			"tx_late_collision", "tx_multicast", "tx_multiple_collision", "tx_packet", "tx_pause", "tx_rx_1024_1518",
			"tx_rx_128_255", "tx_rx_1519_max", "tx_rx_256_511", "tx_rx_512_1023", "tx_rx_64", "tx_rx_65_127",
			"tx_single_collision", "tx_too_long", "tx_too_short", "tx_total_collision"),

		"cpu_flow_control": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "All switch chips have a special port that is called switchX-cpu, this is the CPU port for a " +
				"switch chip, it is meant to forward traffic from a switch chip to the CPU, such a port is required " +
				"for management traffic and for routing features. By default the switch chip ensures that this " +
				"special CPU port is not congested and sends out Pause Frames when link capacity is exceeded to make " +
				"sure the port is not oversaturated, this feature is called CPU Flow Control. Without this feature " +
				"packets that might be crucial for routing or management purposes might get dropped.",
		},
		KeyInvalid: PropInvalidRo,
		"l3_hw_offloading": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Layer 3 Hardware Offloading (L3HW, otherwise known as IP switching or HW routing) allows to " +
				"offload some router features onto the switch chip. This allows reaching wire speeds when routing " +
				"packets, which simply would not be possible with the CPU.",
		},
		// "mirror_egress_target": {
		// 	Type:     schema.TypeString,
		// 	Optional: true,
		// 	Default:  "none",
		// 	Description: "Selects a single mirroring egress target port, only available on 88E6393X, 88E6191X and " +
		// 		"88E6190 switch chips. Mirrored packets from mirror-egress (see the property in port menu) will be " +
		// 		"sent to the selected port.",
		// },
		"mirror_source": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects a single mirroring source port. Ingress and egress traffic will be sent to the " +
				"mirror-target port. Note that mirror-target port has to belong to the same switch (see which port " +
				"belongs to which switch in /interface ethernet menu).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirror_target": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects a single mirroring target port. Mirrored packets from mirror-source and mirror " +
				"(see the property in rule and host table) will be sent to the selected port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirror_egress_target": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects a single mirroring egress target port, only available on 88E6393X, 88E6191X and " +
				"88E6190 switch chips. Mirrored packets from `mirror-egress` (see the property in port menu) will be sent " +
				"to the selected port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the switch."),
		"switch_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "*0",
			Description: "Switch-chip id. Default .id = *0",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\*\d+$`),
				"The string must contain an identifier in MikroTik format: '*0'"),
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Switch-chip type.",
		},
	}

	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		metadata := GetMetadata(resSchema)

		res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgGet, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			return nil
		}

		id := (*res)[0].GetID(metadata.IdType)
		d.SetId(id)
		d.Set("switch_id", id)

		return MikrotikResourceDataToTerraform((*res)[0], resSchema, d)
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if d.Id() == "" {
			d.SetId(d.Get("switch_id").(string))
		}
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}
		item[".id"] = d.Id()

		err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return resRead(ctx, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   resRead,
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
