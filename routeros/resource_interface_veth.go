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
		MetaResourcePath: PropResourcePath("/interface/veth"),
		MetaId:           PropId(Id),

		"address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "IP address.",
			ValidateFunc: validation.IsCIDR,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
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
		KeyName:    PropName("Interface name."),
		KeyRunning: PropRunningRo,
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
