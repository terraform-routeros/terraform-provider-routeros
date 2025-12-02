package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".about": "main vrf not suitable for export",
    ".id": "*1",
    "disabled": "false",
    "inactive": "false",
    "label-allocation-policy": "per-prefix",
    "name": "bgp-mpls-vpn-1",
    "route-distinguisher": "123",
    "vrf": "main"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/331612228/routing+bgp#id-/routing/bgp-/routing/bgp/vpn
func ResourceRoutingBgpVpn() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bgp/vpn"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		"export": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with the route export.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"filter_chain": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The name of the routing filter chain that is used to filter prefixes before exporting.",
					},
					"filter_select": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The name of the select filter chain that is used to select prefixes to be exported exporting.",
					},
					"redistribute": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Enable redistribution of specified route types from VRF to VPNv4.",
						ValidateFunc: validation.StringInSlice([]string{"bgp", "connected", "dhcp", "fantasy", "modem", "ospf", "rip",
							"static", "vpn"}, false),
					},
					"route_targets": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of route targets added when exporting VPNv4 routes. The accepted RT format is similar " +
							"to the one for Route Distinguishers.",
					},
				},
			},
		},
		"import": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with the route import.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"filter_chain": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The name of the routing filter chain that is used to filter prefixes during import.",
					},
					"route_targets": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of route targets that will be used to import VPNv4 routes. The accepted RT format is " +
							"similar to the one for Route Distinguishers.",
					},
					"router_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The router ID of the BGP instance that will be used for the BGP best path selection algorithm.",
					},
				},
			},
		},
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the instance this VPN is assigned to.",
		},
		"label_allocation_policy": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Label allocationpolicy.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"per-prefix", "per-vrf"}, false),
		},
		KeyName: PropName("VPN instance name."),
		"route_distinguisher": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Helps to distinguish between overlapping routes from multiple VRFs. Should be unique " +
				"per VRF. Accepts 3 types of formats.",
		},
		KeyVrf: PropVrfRw,
	}

	return &schema.Resource{
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
