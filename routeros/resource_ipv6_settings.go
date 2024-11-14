package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "accept-redirects": "yes-if-forwarding-disabled",
  "accept-router-advertisements": "yes-if-forwarding-disabled",
  "disable-ipv6": "false",
  "forward": "true",
  "max-neighbor-entries": "8192"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841817/IP+Settings#IPSettings-IPv6Settings
func ResourceIpv6Settings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/settings"),
		MetaId:           PropId(Id),

		"accept_redirects": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether to accept ICMP redirect messages. Typically should be enabled on the host and disabled " +
				"on routers.",
			ValidateFunc:     validation.StringInSlice([]string{"no", "yes-if-forwarding-disabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"accept_router_advertisements": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Accept router advertisement (RA) messages. If enabled, the router will be able to get the " +
				"address using stateless address configuration.",
			ValidateFunc:     validation.StringInSlice([]string{"no", "yes", "yes-if-forwarding-disabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"disable_ipv6": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable/disable system wide IPv6 settings (prevents LL address generation).",
		},
		"forward": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable/disable packet forwarding between interfaces.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_neighbor_entries": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "A maximum number or IPv6 neighbors. Since RouterOS version 7.1, the default value depends " +
				"on the installed amount of RAM. It is possible to set a higher value than the default, but it increases " +
				"the risk of out-of-memory condition. The default values for certain RAM sizes:\n  * 1024 for 64 MB,\n  * 2048 " +
				"for 128 MB,\n  * 4096 for 256 MB,\n  * 8192 for 512 MB,\n  * 16384 for 1024 MB or higher.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multipath_hash_policy": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IPv6 Hash policy used for ECMP routing in `/ipv6/settings` menu\n  * l3 -- layer-3 hashing of src " +
				"IP, dst IP, flow label, IP protocol\n  * l3-inner -- layer-3 hashing or inner layer-3 hashing if available" +
				"\n  * l4 -- layer-4 hashing of src IP, dst IP, IP protocol, src port, dst port.",
			ValidateFunc:     validation.StringInSlice([]string{"l3", "l4", "l3-inner"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
