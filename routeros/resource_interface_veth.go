package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*8",
    "address": "192.168.100.2/24",
    "comment": "comment",
    "disabled": "false",
    "gateway": "192.168.100.1",
    "name": "veth1",
    "running": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Container
func ResourceInterfaceVeth() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/interface/veth"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("mac_address"),

		"address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Ip address.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dhcp": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"gateway": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Gateway IP address.",
			ValidateFunc: validation.IsIPv4Address,
		},
		"gateway6": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Gateway IPv6 address.",
			ValidateFunc: validation.IsIPv6Address,
		},
		KeyMacAddress: PropMacAddressRw("MAC address.", false),
		KeyName:       PropName("Interface name."),
		KeyRunning:    PropRunningRo,
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
