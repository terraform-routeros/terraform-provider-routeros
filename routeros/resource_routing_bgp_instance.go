package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*2",
    "as": "65535",
    "disabled": "false",
    "inactive": "false",
    "name": "bgp-instance-1"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/
func ResourceRoutingBgpInstance() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bgp/instance"),
		MetaId:           PropId(Id),

		"as": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "32-bit BGP autonomous system number. Value can be entered in AS-Plain and AS-Dot formats. " +
				"The parameter is also used to set up the BGP confederation, in the following format: confederation_as/as. " +
				"For example, if your AS is 34 and your confederation AS is 43, then as configuration should be as=43/34.",
		},
		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "In case this instance is a route reflector: the cluster ID of the router reflector cluster " +
				"to this instance belongs. This attribute helps to recognize routing updates that come from another route " +
				"reflector in this cluster and avoid routing information looping. Note that normally there is only one " +
				"route reflector in a cluster; in this case, `cluster-id` does not need to be configured and BGP router " +
				"ID is used instead.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"ignore_as_path_len": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to ignore the `AS_PATH` attribute in the BGP route selection algorithm. Works on input.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Instance name."),
		"router_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "BGP Router ID to be used. Use the ID from the `/routing/router-id` configuration by specifying " +
				"the reference name, or set the ID directly by specifying IP.Equal router-ids are also used to group " +
				"peers into one instance.",
		},
		"routing_table": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the routing table, to install routes in.",
		},
		"vrf": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the VRF BGP connections operates on. By default always use the `main` routing table.",
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
