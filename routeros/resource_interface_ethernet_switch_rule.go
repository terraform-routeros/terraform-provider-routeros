package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
???
*/

// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features#SwitchChipFeatures-RuleTable
func ResourceInterfaceEthernetSwitchRule() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/rule"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInvalid:  PropInvalidRo,

		"copy_to_cpu": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send a frame copy to switch CPU port from a frame with matching MAC destination address " +
				"(matching destination or source address for CRS3xx series switches).",
		},
		"dst_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching destination IP address and mask.",
		},
		"dst_address6": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching destination IPv6 address and mask.",
		},
		"dst_mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching destination MAC address and mask.",
		},
		"dst_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching destination protocol port number or range.",
			ValidateFunc: Validation64k,
		},
		"dscp": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching DSCP field of the packet.",
			ValidateFunc: validation.IntBetween(0, 63),
		},
		"flow_label": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching IPv6 flow label.",
			ValidateFunc: validation.IntBetween(0, 1048575),
		},
		"mac_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching particular MAC protocol specified by protocol name or number (skips VLAN tags if any).",
		},
		"mirror": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send a frame copy to mirror-target port from a frame with matching MAC destination address " +
				"(matching destination or source address for CRS3xx series switches).",
		},
		"mirror_ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Selects multiple mirroring target ports, only available on 88E6393X switch chip. " +
				"Matched packets in the ACL rule will be copied and sent to selected ports.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: AlwaysPresentNotUserProvided,
			}},
		"new_dst_ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Changes the destination port as specified, multiple ports allowed, including a switch CPU port. An empty " +
				"setting will drop the packet. When the parameter is not used, the packet will be accepted.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: AlwaysPresentNotUserProvided,
			},
		},
		"new_vlan_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Changes the VLAN ID to the specified value or adds a new VLAN tag if one was not already present " +
				"(the property only applies to the Atheros8316, and 88E6393X switch chips).",
			ValidateFunc:     validation.IntBetween(0, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"new_vlan_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Changes the VLAN priority field (priority code point, the property only applies to Atheros8327, QCA8337 " +
				"and Atheros8316 switch chips).",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ports": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "Name of the interface on which the rule will apply on the received traffic, multiple ports are allowed.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: AlwaysPresentNotUserProvided,
			},
		},
		"protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching particular IP protocol specified by protocol name or number.",
		},
		"rate": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Sets ingress traffic limitation (bits per second) for matched traffic, can only be applied to the first 32 rule slots " +
				"(the property only applies to Atheros8327/QCA8337 switch chips).",
			ValidateFunc:     validation.IntAtLeast(0),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"redirect_to_cpu": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Changes the destination port of a matching packet to the switch CPU.",
		},
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching source IP address and mask.",
		},
		"src_address6": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching source IPv6 address and mask.",
		},
		"src_mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matching source MAC address and mask.",
		},
		"src_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching source protocol port number or range.",
			ValidateFunc: Validation64k,
		},
		"switch": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Matching switch group on which will the rule apply.",
		},
		"traffic_class": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching IPv6 traffic class.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"vlan_header": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matching VLAN header, whether the VLAN header is present or not (the property only applies to the " +
				"Atheros8316, Atheros8327, QCA8337, 88E6393X switch chips).",
			ValidateFunc: validation.StringInSlice([]string{"not-present", "present"}, false),
		},
		"vlan_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Matching VLAN ID (the property only applies to the Atheros8316, Atheros8327, QCA8337, 88E6393X switch chips).",
			ValidateFunc:     validation.IntBetween(0, 4095),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matching VLAN priority (priority code point).",
			ValidateFunc: validation.IntBetween(0, 7),
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
