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

// https://help.mikrotik.com/docs/display/ROS/Neighbor+discovery
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
		"discover_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "An option to adjust the frequency at which neighbor discovery packets are transmitted. " +
				"The setting is available since RouterOS version 7.16.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_mac_phy_config": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send MAC/PHY Configuration/Status TLV in LLDP, which indicates the interface " +
				"capabilities, current setting of the duplex status, bit rate, and auto-negotiation. Only applies " +
				"to the Ethernet interfaces. While TLV is optional in LLDP, it is mandatory when sending LLDP-MED, " +
				"meaning this TLV will be included when necessary even though the property is configured as disabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_max_frame_size": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send Maximum Frame Size TLV in LLDP, which indicates the maximum frame size capability" +
				" of the interface in bytes (`l2mtu + 18`). Only applies to the Ethernet interfaces.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_med_net_policy_vlan": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Advertised VLAN ID for LLDP-MED Network Policy TLV. This allows assigning a VLAN ID for " +
				"LLDP-MED capable devices, such as VoIP phones. The TLV will only be added to interfaces where LLDP-MED " +
				"capable devices are discovered. Other TLV values are predefined and cannot be changed:" +
				"\n  * Application Type - Voice" +
				"\n  * VLAN Type - Tagged" +
				"\n  * L2 Priority - 0" +
				"\n  * DSCP Priority - 0\n" +
				"When used together with the bridge interface, the (R/M)STP protocol should be enabled with protocol-mode setting.\n" +
				"Additionally, other neighbor discovery protocols (e.g. CDP) should be excluded using protocol setting to " +
				"avoid LLDP-MED misconfiguration.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_poe_power": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Two specific TLVs facilitate Power over Ethernet (PoE) management between Power Sourcing " +
				"Equipment (PSE) and Powered Devices (PD).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lldp_vlan_info": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "An option whether to send IEEE 802.1 Organizationally Specific TLVs in LLDP related to VLANs. " +
				"The setting is available since RouterOS version 7.16.",
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
