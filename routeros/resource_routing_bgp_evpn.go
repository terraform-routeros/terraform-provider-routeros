package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
[
  {
    ".about": "rd must be specified for bundled vlans",
    ".id": "*1",
    "export.route-targets": "1010:1010,1010:1020",
    "instance": "test",
    "name": "test"
  }
]
*/

// https://help.mikrotik.com/docs/display/ROS/
func ResourceRoutingBgpEvpn() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bgp/evpn"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "BGP instance this EVPN is assigned to.",
		},
		"export": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A group of parameters associated with the route export.",
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"route_targets": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of route targets that will be added to EVPN routes when exporting.",
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
					"route_targets": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "List of route targets that will be used to import EVPN routes.",
					},
				},
			},
		},
		KeyName: PropName("Name of the entry."),
		"rd": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the value that gets attached to route so that receiving routers can distinguish " +
				"advertisements that may otherwise look the same. Used to distinguish between tenants using overlapping " +
				"IP ranges. Also can be used to simplify convergence and redundancy within Virtual Network. RDs form " +
				"MLAG pairs should be unique, too.",
		},
		"vni": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Range of Virtual Network Identifiers.",
		},
		"vrf": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VRF name.",
		},
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
