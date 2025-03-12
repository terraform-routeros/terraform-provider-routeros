package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "accept-redirects": "yes-if-forwarding-disabled",
  "accept-router-advertisements": "yes-if-forwarding-disabled",
  "allow-fast-path": "true",
  "disable-ipv6": "false",
  "disable-link-local-address": "false",
  "forward": "true",
  "ipv6-fast-path-active": "true",
  "ipv6-fast-path-bytes": "0",
  "ipv6-fast-path-packets": "0",
  "ipv6-fasttrack-active": "false",
  "ipv6-fasttrack-bytes": "0",
  "ipv6-fasttrack-packets": "0",
  "max-neighbor-entries": "16384",
  "min-neighbor-entries": "4096",
  "multipath-hash-policy": "l3",
  "soft-max-neighbor-entries": "8192",
  "stale-neighbor-detect-interval": "30",
  "stale-neighbor-timeout": "60"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841817/IP+Settings#IPSettings-IPv6Settings
func ResourceIpv6Settings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/settings"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("ipv6_fast_path_bytes", "ipv6_fast_path_packets", "ipv6_fasttrack_bytes", "ipv6_fasttrack_packets"),

		"allow_fast_path": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Allows Fast Path.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
		"disable_link_local_address": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Disable automatic link-local address generation for non-VPN interfaces. This can be used when " +
				"manually configured link-local addresses are being used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"forward": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enable/disable packet forwarding between interfaces.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipv6_fast_path_active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether fast-path is active.",
		},
		"ipv6_fasttrack_active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether fasttrack is active.",
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
		"min_neighbor_entries": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Minimal number of IPv6/Neighbor entries, for which device must allocate memory.",
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
		"soft_max_neighbor_entries": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Expected maximum number of IPv6/Neighbor entries which system should handle.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"stale_neighbor_detect_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: TimeEqual,
		},
		"stale_neighbor_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Timeout after which stale IPv6/Neighbor entries should be purged.",
			DiffSuppressFunc: TimeEqual,
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
