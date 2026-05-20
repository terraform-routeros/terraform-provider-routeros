package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
	"bridge":"bridge1",
	"peer-port":"stack-link"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Multi-chassis+Link+Aggregation+Group
func ResourceInterfaceBridgeMlag() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/mlag"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface where MLAG is being created.",
		},
		"heartbeat": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This setting controls how often heartbeat messages are sent to check the connection between peers. " +
				"If no heartbeat message is received for three intervals in a row, the peer logs a warning about " +
				"potential communication problems. If set to none, heartbeat messages are not sent at all.",
			DiffSuppressFunc: TimeEqual,
		},
		"peer_port": {
			Type:     schema.TypeString,
			Required: true,
			Description: "An interface that will be used as a peer port. Both peer devices are using inter-chassis " +
				"communication over these peer ports to establish MLAG and update the host table. Peer port should be " +
				"isolated on a different untagged VLAN using a pvid setting. Peer port can be configured as a bonding " +
				"interface.",
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "This setting changes the priority for selecting the primary MLAG node. A lower number means " +
				"higher priority. If both MLAG nodes have the same priority, the one with the lowest bridge MAC address " +
				"will become the primary device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	validateLegacyMlag := func(*schema.ResourceData) diag.Diagnostics {
		atLeast722, err := routerOSVersionAtLeast("7.22")
		if err != nil {
			return diag.FromErr(err)
		}

		if atLeast722 {
			return diag.Errorf("routeros_bridge_mlag is only supported before RouterOS 7.22; manage MLAG with routeros_interface_bridge mlag_* attributes on RouterOS 7.22 and newer")
		}

		return nil
	}

	readLegacyMlag := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		diags := DefaultSystemRead(resSchema)(ctx, d, m)
		if diags.HasError() {
			return diags
		}

		atLeast722, err := routerOSVersionAtLeast("7.22")
		if err != nil {
			return diag.FromErr(err)
		}

		if atLeast722 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "routeros_bridge_mlag is legacy on RouterOS 7.22 and newer",
				Detail:   "Manage MLAG through routeros_interface_bridge mlag_* attributes and migrate off routeros_bridge_mlag on RouterOS 7.22 and newer.",
			})
		}

		return diags
	}

	return &schema.Resource{
		Description: "Legacy MLAG resource for RouterOS versions before 7.22. " +
			"Use the `mlag_heartbeat`, `mlag_peer_port`, and `mlag_priority` attributes on `routeros_interface_bridge` for RouterOS 7.22 and newer.",
		CreateContext: DefaultValidateCreate(resSchema, validateLegacyMlag),
		ReadContext:   readLegacyMlag,
		UpdateContext: DefaultValidateUpdate(resSchema, validateLegacyMlag),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
