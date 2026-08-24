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
  ".id": "*5",
  "default": "false",
  "ecn": "false",
  "ecn-actual": "false",
  "hw-offloaded": "true",
  "inactive": "false",
  "queue-buffers": "auto",
  "schedule": "high-priority-group",
  "shared-pool-index": "0",
  "traffic-class": "5",
  "tx-manager": "default",
  "use-shared-buffers": "true",
  "weight": "8",
  "wred": "false",
  "wred-actual": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CRS3xx,+CRS5xx,+CCR2116,+CCR2216+QoS
// ResourceInterfaceEthernetSwitchQosTxManagerQueue configures the egress
// scheduler for one traffic class within a tx-manager. The eight queues (one per
// traffic class 0..7) are auto-created by the chip, so this resource adopts an
// existing queue keyed by `traffic_class` (within `tx_manager`) and can only
// modify it, never add or delete it.
func ResourceInterfaceEthernetSwitchQosTxManagerQueue() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/tx-manager/queue"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("hw_offloaded", "inactive", "ecn_actual", "wred_actual"),

		"traffic_class": {
			Type:         schema.TypeInt,
			Required:     true,
			ForceNew:     true,
			Description:  "Traffic class (egress queue index, 0..7) this queue configures.",
			ValidateFunc: validation.IntBetween(0, 7),
		},
		"tx_manager": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Default:     "default",
			Description: "Name of the tx-manager (egress scheduler) this queue belongs to.",
		},
		"schedule": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Scheduling group for this queue.",
			ValidateFunc:     validation.StringInSlice([]string{"high-priority-group", "low-priority-group", "strict-priority"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"weight": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Weight of this queue within its scheduling group (ignored for `strict-priority`).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"queue_buffers": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Amount of buffers dedicated to this queue (`auto` or a numeric value).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_shared_buffers": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether this queue may use the shared buffer pool.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"shared_pool_index": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Index of the shared buffer pool used when `use_shared_buffers` is enabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ecn": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable Explicit Congestion Notification (ECN) marking for this queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wred": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable Weighted Random Early Detection (WRED) for this queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether this queue is the default queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		filter := []string{
			fmt.Sprintf("traffic-class=%d", d.Get("traffic_class").(int)),
			"tx-manager=" + d.Get("tx_manager").(string),
		}
		res, err := ReadItemsFiltered(filter, metadata.Path, m.(Client))
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		if len(*res) == 0 {
			d.SetId("")
			return diag.FromErr(errorNoLongerExists)
		}
		if len(*res) > 1 {
			return diag.FromErr(fmt.Errorf("more than one tx-manager queue found for traffic-class=%d tx-manager=%s",
				d.Get("traffic_class").(int), d.Get("tx_manager").(string)))
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		// These identify the (auto-created) queue row and are read-only on it.
		delete(item, "traffic-class")
		delete(item, "tx-manager")

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
