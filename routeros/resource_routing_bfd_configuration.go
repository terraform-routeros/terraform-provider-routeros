package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    ".nextid": "*2",
    "address-list": "",
    "addresses": "0.0.0.0/0,0.0.0.0/0",
    "disabled": "false",
    "forbid-bfd": "false",
    "inactive": "false",
    "interfaces": "lo,lo",
    "min-rx": "200ms",
    "min-tx": "200ms",
    "multiplier": "5",
    "vrf": "main"
  },
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/331612210/routing+bfd
// https://help.mikrotik.com/docs/spaces/ROS/pages/191299691/BFD
func ResourceRoutingBfdConfiguration() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bfd/configuration"),
		MetaId:           PropId(Id),

		"address_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the address list in which users IP address will be added.",
		},
		"addresses": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
			Description: "Set of IP (v4 or v6) addresses or CIDR networks.",
		},
		KeyDisabled: PropDisabledRw,
		"forbid_bfd": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "",
		},
		KeyInactive: PropInactiveRo,
		"interfaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of interfaces.",
		},
		"min_rx": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: TimeEqual,
		},
		"min_tx": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: TimeEqual,
		},
		"multiplier": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "",
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
