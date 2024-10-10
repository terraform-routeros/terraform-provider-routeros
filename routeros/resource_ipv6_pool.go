package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*2",
    "dynamic": "false",
    "name": "test-pool",
    "prefix": "2001:db8:12::/48",
    "prefix-length": "64"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IP+Pools#IPPools-IPv6Pool
func ResourceIpv6Pool() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/pool"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("expire_time"),

		KeyDynamic: PropDynamicRo,
		KeyName:    PropName("Descriptive name of the pool."),
		"prefix": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Ipv6 address prefix.",
			ValidateFunc:     validation.IsCIDR,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"prefix_length": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "The option represents the prefix size that will be given out to the client.",
			ValidateFunc: validation.IntBetween(1, 128),
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
