package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*2",
  "invalid": "false",
  "map": "default",
  "name": "sfp-sfpplus1",
  "pfc": "disabled",
  "profile": "5-Standard (BE/Default)",
  "running": "true",
  "switch": "switch1",
  "trust-l2": "trust",
  "trust-l3": "trust",
  "tx-manager": "default"
  // ... plus many read-only per-queue byte/packet counters and caps
}
*/

// https://help.mikrotik.com/docs/display/ROS/CRS3xx,+CRS5xx,+CCR2116,+CCR2216+QoS
// ResourceInterfaceEthernetSwitchQosPort configures per-port QoS on the switch
// chip: which default profile to use and whether to trust the incoming L2 (PCP)
// and/or L3 (DSCP) markings. Port rows are auto-created by the chip (one per
// switch port), so this resource adopts an existing row keyed by `name` and can
// only modify it, never add or delete it.
func ResourceInterfaceEthernetSwitchQosPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/port"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("pfc_rx", "pfc_unknown",
			"tx_byte", "tx_packet", "drop_byte", "drop_packet",
			"queue0_byte_cap", "queue0_packet_cap", "queue0_shared_byte_cap", "queue0_shared_packet_cap",
			"queue1_byte_cap", "queue1_packet_cap", "queue1_shared_byte_cap", "queue1_shared_packet_cap",
			"queue2_byte_cap", "queue2_packet_cap", "queue2_shared_byte_cap", "queue2_shared_packet_cap",
			"queue3_byte_cap", "queue3_packet_cap", "queue3_shared_byte_cap", "queue3_shared_packet_cap",
			"queue4_byte_cap", "queue4_packet_cap", "queue4_shared_byte_cap", "queue4_shared_packet_cap",
			"queue5_byte_cap", "queue5_packet_cap", "queue5_shared_byte_cap", "queue5_shared_packet_cap",
			"queue6_byte_cap", "queue6_packet_cap", "queue6_shared_byte_cap", "queue6_shared_packet_cap",
			"queue7_byte_cap", "queue7_packet_cap", "queue7_shared_byte_cap", "queue7_shared_packet_cap",
			"tx_queue0_byte", "tx_queue0_packet", "tx_queue1_byte", "tx_queue1_packet",
			"tx_queue2_byte", "tx_queue2_packet", "tx_queue3_byte", "tx_queue3_packet",
			"tx_queue4_byte", "tx_queue4_packet", "tx_queue5_byte", "tx_queue5_packet",
			"tx_queue6_byte", "tx_queue6_packet", "tx_queue7_byte", "tx_queue7_packet",
			"drop_queue0_byte", "drop_queue0_packet", "drop_queue1_byte", "drop_queue1_packet",
			"drop_queue2_byte", "drop_queue2_packet", "drop_queue3_byte", "drop_queue3_packet",
			"drop_queue4_byte", "drop_queue4_packet", "drop_queue5_byte", "drop_queue5_packet",
			"drop_queue6_byte", "drop_queue6_packet", "drop_queue7_byte", "drop_queue7_packet",
		),

		KeyName:    PropName("Name of the switch port to configure."),
		KeyInvalid: PropInvalidRo,
		KeyRunning: PropRunningRo,
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Default `qos profile` applied to traffic on this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"trust_l2": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to trust the L2 (802.1p PCP) priority of ingress traffic on this port.",
			ValidateFunc:     validation.StringInSlice([]string{"trust", "ignore", "keep"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"trust_l3": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to trust the L3 (IP DSCP) priority of ingress traffic on this port.",
			ValidateFunc:     validation.StringInSlice([]string{"trust", "ignore", "keep"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"map": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the QoS map set used for classification on this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pfc": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Priority Flow Control (PFC) mode for this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_manager": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the `qos tx-manager` (egress scheduler) applied to this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"switch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the switch this port belongs to.",
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItems(&ItemId{Name, d.Get("name").(string)}, metadata.Path, m.(Client))
		if err != nil {
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

		// name identifies the port row and is read-only on it.
		delete(item, "name")

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
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
