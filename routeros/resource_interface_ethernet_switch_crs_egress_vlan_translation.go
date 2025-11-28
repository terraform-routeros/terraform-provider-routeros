package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  ".nextid": "*FFFFFFFF",
  "comment": "",
  "customer-vid": "200",
  "customer-vlan-format": "any",
  "disabled": "false",
  "dynamic": "false",
  "new-customer-vid": "0",
  "pcp-propagation": "false",
  "ports": "ether1",
  "service-vlan-format": "any"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841835/CRS1xx+and+2xx+series+switches#CRS1xxand2xxseriesswitches-Ingress%2FEgressVLANTranslation
func ResourceInterfaceEthernetSwitchCrsEgressVlanTranslation() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/egress-vlan-translation"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"customer_dei": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching DEI of the customer tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"customer_pcp": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching PCP of the customer tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"customer_vid": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching the VLAN ID of the customer tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"customer_vlan_format": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of frames with customer tag for which VLAN translation rule is valid.",
			ValidateFunc:     validation.StringInSlice([]string{"any", "priority-tagged-or-tagged", "tagged", "untagged-or-tagged"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"new_customer_vid": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The new customer VLAN ID replaces the matching customer VLAN ID. If set to 4095 and ingress " +
				"VLAN translation is used, then traffic is dropped.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"new_service_vid": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The new service VLAN ID replaces the matching service VLAN ID.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pcp_propagation": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Enables or disables PCP propagation.If the port type is Edge, the customer PCP is copied from " +
				"the service PCP.If the port type is Network, the service PCP is copied from the customer PCP.",
		},
		"ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Matching switch ports for VLAN translation rule.",
		},
		"service_dei": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching DEI of the service tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"service_pcp": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching PCP of the service tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"service_vid": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matching VLAN ID of the service tag.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"service_vlan_format": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of frames with service tag for which VLAN translation rule is valid.",
			ValidateFunc:     validation.StringInSlice([]string{"any", "priority-tagged-or-tagged", "tagged", "untagged-or-tagged"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"swap_vids": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
	}

	return &schema.Resource{
		Description: "Resource for managing CRS (Cloud Router Switch) series device properties.",

		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
