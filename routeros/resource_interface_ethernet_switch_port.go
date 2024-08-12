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
    ".id": "*1000000",
     "default-vlan-id": "0",
     "invalid": "false",
     "name": "switch1-cpu",
    "rx-1024-1518": "0",
    "rx-128-255": "0",
    "rx-1519-max": "0",
    "rx-256-511": "0",
    "rx-512-1023": "0",
    "rx-64": "0",
    "rx-65-127": "0",
    "rx-align-error": "0",
    "rx-broadcast": "0",
    "rx-bytes": "0",
    "rx-fcs-error": "0",
    "rx-fragment": "0",
    "rx-multicast": "0",
    "rx-overflow": "0",
    "rx-pause": "0",
    "rx-too-long": "0",
    "rx-too-short": "0",
     "switch": "switch1",
    "tx-1024-1518": "0",
    "tx-128-255": "0",
    "tx-1519-max": "0",
    "tx-256-511": "0",
    "tx-512-1023": "0",
    "tx-64": "0",
    "tx-65-127": "0",
    "tx-broadcast": "0",
    "tx-bytes": "0",
    "tx-collision": "0",
    "tx-deferred": "0",
    "tx-excessive-collision": "0",
    "tx-excessive-deferred": "0",
    "tx-late-collision": "0",
    "tx-multicast": "0",
    "tx-multiple-collision": "0",
    "tx-pause": "0",
    "tx-single-collision": "0",
    "tx-too-long": "0",
    "tx-underrun": "0",
     "vlan-header": "leave-as-is",
     "vlan-mode": "disabled"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/L3+Hardware+Offloading#L3HardwareOffloading-SwitchPortConfiguration
// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features
func ResourceInterfaceEthernetSwitchPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/port"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("name", "rx_1024_1518", "rx_128_255", "rx_1519_max", "rx_256_511", "rx_512_1023", "rx_64",
			"rx_65_127", "rx_align_error", "rx_broadcast", "rx_bytes", "rx_fcs_error", "rx_fragment", "rx_multicast", "rx_overflow",
			"rx_pause", "rx_too_long", "rx_too_short", "tx_1024_1518", "tx_128_255", "tx_1519_max", "tx_256_511", "tx_512_1023", "tx_64",
			"tx_65_127", "tx_broadcast", "tx_bytes", "tx_collision", "tx_deferred", "tx_excessive_collision", "tx_excessive_deferred",
			"tx_late_collision", "tx_multicast", "tx_multiple_collision", "tx_pause", "tx_single_collision", "tx_too_long", "tx_underrun",
			"driver_tx_byte", "driver_rx_packet", "driver_rx_byte", "driver_tx_packet", "tx_carrier_sense_error",
			"rx_jabber", "tx_rx_65_127", "tx_rx_512_1023", "rx_unicast", "tx_fcs_error", "tx_rx_128_255", "tx_unicast",
			"tx_rx_1024_max", "tx_rx_256_511", "rx_error_events", "tx_rx_64",
		),

		"default_vlan_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adds a VLAN tag with the specified VLAN ID on all untagged ingress traffic on a port, should be used " +
				"with ```vlan-header``` set to ```always-strip``` on a port to configure the port to be the access port. For hybrid ports " +
				"```default-vlan-id``` is used to tag untagged traffic. If two ports have the same ```default-vlan-id```, then VLAN tag is " +
				"not added since the switch chip assumes that traffic is being forwarded between access ports.",
			ValidateFunc:     validation.StringMatch(regexp.MustCompile(`auto|\d+`), `Value must be "auto" or integer: 0..4095`),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInvalid: PropInvalidRo,
		"mirror_egress": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send egress packet copy to the `mirror-egress-target` port, only available on " +
				"88E6393X, 88E6191X and 88E6190 switch chips.",
		},
		"mirror_ingress": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send ingress packet copy to the `mirror-ingress-target` port, only available on " +
				"88E6393X, 88E6191X and 88E6190 switch chips.",
		},
		"mirror_ingress_target": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects a single mirroring ingress target port, only available on  88E6393X, 88E6191X and " +
				"88E6190 switch chips. Mirrored packets from `mirror-ingress` will be sent to the selected port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName:    PropName("Port name."),
		KeyRunning: PropRunningRo,
		"switch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the switch.",
		},
		"vlan_header": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets action which is performed on the port for egress traffic.",
			ValidateFunc:     validation.StringInSlice([]string{"add-if-missing", "always-strip", "leave-as-is"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Changes the VLAN lookup mechanism against the VLAN Table for ingress traffic.",
			ValidateFunc:     validation.StringInSlice([]string{"check", "disabled", "fallback", "secure"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItems(&ItemId{Name, d.Get("name").(string)}, metadata.Path, m.(Client))
		if err != nil {
			// API/REST client error.
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(errorNoLongerExists)
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}

		err = m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
