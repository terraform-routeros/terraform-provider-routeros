package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
[
  {
    ".id": "*1",
    "bridge": "bridge",
    "comment": "Management",
    "current-tagged": "bridge,ether2,ether3",
    "current-untagged": "",
    "disabled": "false",
    "dynamic": "false",
    "tagged": "ether2,ether4,ether5,bridge,ether3",
    "untagged": "",
    "vlan-ids": "2"
  },
  {...}
]
*/

// ResourceInterfaceBridgeVlan https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge#Bridge_VLAN_Filtering
func ResourceInterfaceBridgeVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/vlan"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface which the respective VLAN entry is intended for.",
		},
		KeyComment: PropCommentRw,
		"current_tagged": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"current_untagged": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"tagged": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interface list with a VLAN tag adding action in egress. This setting accepts comma " +
				"separated values. E.g. tagged=ether1,ether2.",
		},
		"untagged": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interface list with a VLAN tag removing action in egress. This setting accepts comma " +
				"separated values. E.g. untagged=ether3,ether4",
		},
		"vlan_ids": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The list of VLAN IDs for certain port configuration. This setting accepts VLAN ID range " +
				"as well as comma separated values. E.g. vlan-ids=100-115,120,122,128-130.",
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
