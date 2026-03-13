package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "comment": "",
    "disabled": "false",
    "dynamic": "false",
    "dynamic-id": "",
    "id": "10.10.10.10",
    "inactive": "false",
    "name": "router-id-1",
    "select-dynamic-id": "any",
    "select-from-vrf": "main"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/30474294/routing+id
func ResourceRoutingId() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/id"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("router_id:id"),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"dynamic_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Currently selected ID.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Reference name."),
		"router_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsIPAddress,
			Description: "Parameter to explicitly set the Router ID. If not specified, it can be elected from one " +
				"of the configured IP addresses on the router.",
		},
		"select_dynamic_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "States what IP addresses to use for ID election.",
			ValidateFunc: validation.StringInSlice([]string{
				"any", "lowest", "only-active", "only-loopback", "only-static", "only-vrf",
			}, false),
		},
		"select_from_vrf": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VRF from which to select IP addresses for the ID election.",
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
