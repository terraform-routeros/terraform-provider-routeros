package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "discover-interface-list": "LAN",
  "lldp-med-net-policy-vlan": "disabled",
  "mode": "tx-and-rx",
  "protocol": "cdp,lldp,mndp"
}
*/

// https://help.mikrotik.com/docs/display/ROS/MAC+server
func ResourceIpNeighborDiscoverySettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/neighbor/discovery-settings"),
		MetaId:           PropId(Id),

		"discover_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interface list on which members the discovery protocol will run on.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_med_net_policy_vlan": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Advertised VLAN ID for LLDP-MED Network Policy TLV. This allows assigning a VLAN ID for 
			LLDP-MED capable devices, such as VoIP phones. The TLV will only be added to interfaces where LLDP-MED 
			capable devices are discovered. Other TLV values are predefined and cannot be changed:

			- Application Type - Voice
			- VLAN Type - Tagged
			- L2 Priority - 0
			- DSCP Priority - 0
		
		When used together with the bridge interface, the (R/M)STP protocol should be enabled with protocol-mode setting. 
		
		Additionally, other neighbor discovery protocols (e.g. CDP) should be excluded using protocol setting to 
		avoid LLDP-MED misconfiguration.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects the neighbor discovery packet sending and receiving mode. The setting is " +
				"available since RouterOS version 7.7.",
			ValidateFunc:     validation.StringInSlice([]string{"rx-only", "tx-only", "tx-and-rx"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"protocol": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "List of used discovery protocols.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cdp", "lldp", "mndp"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
