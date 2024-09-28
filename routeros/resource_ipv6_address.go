package routeros

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
[
  {
    ".id": "*1",
    "actual-interface": "ether1",
    "address": "fe80::5605:abff:fecd:1231/64",
    "advertise": "false",
    "disabled": "false",
    "dynamic": "true",
    "eui-64": "false",
    "from-pool": "",
    "interface": "ether1",
    "invalid": "false",
    "link-local": "true",
    "no-dad": "false"
  }
]
*/

// ResourceIPv6Address https://help.mikrotik.com/docs/display/ROS/IP+Addressing
func ResourceIPv6Address() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/address"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "IPv6 address. Using the eui_64 and from_pool options can transform the original address! " +
				"[See docs](https://wiki.mikrotik.com/wiki/Manual:IPv6/Address#Properties)",
			AtLeastOneOf: []string{"address", "from_pool"},
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// This check is very dirty, be careful!
				/* eui_64 == true or from_pool != ""

				  After applying this test step, the plan was not empty.

				  # routeros_ipv6_address.test_v6_address will be updated in-place
				  ~ resource "routeros_ipv6_address" "test_v6_address" {
					~ address          = "fc00:3::5c30:77ff:fe61:33ac/64" -> "fc00:3::/64"
					id               = "*1D"
					# (9 unchanged attributes hidden)
				}
				*/

				if old == new {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				// k = address, old = fc00:3::5c30:77ff:fe61:33ac/64, new = fc00:3::/64

				// eui_64: true, address: "fc00:3::/64" ===> "fc00:3::/64" -> "fc00:3::5c30:77ff:fe61:33ac/64"
				addr := strings.SplitN(new, "/", 2)
				if len(addr) == 2 {
					addr[1] = "/" + addr[1]
				}

				if len(addr) == 2 && len(old) >= len(new) {
					if old[:len(addr[0])] == addr[0] && old[len(old)-len(addr[1]):] == addr[1] {
						return true
					}
				}

				/*
					/ipv6/pool/print
					# NAME   PREFIX       PREFIX-LENGTH
					0 pool1  fc00:3::/62             64
				*/

				// from_pool: pool1, address: "::1/64" ===>  "::1/64" -> "fc00:3::1/64"
				if len(old) >= len(new) && strings.HasSuffix(old, new) {
					return true
				}

				// eui_64: true, from_pool: pool1, address: "::1/64" ===> "::1/64" -> "fc00:3::5c30:77ff:fe61:33ac/64"
				// N.B. We consider this a configuration error until such a configuration is encountered by someone.
				// eui_64: true and from_pool: pool1 ===> "" -> ""
				return false
			},
		},
		"advertise": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to enable stateless address configuration. The prefix of that address is " +
				"automatically advertised to hosts using ICMPv6 protocol. The option is set by default for addresses " +
				"with prefix length 64.",
		},
		"actual_interface": { // RO
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the actual interface the logical one is bound to.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"deprecated": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether address is deprecated",
		},
		"eui_64": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to calculate EUI-64 address and use it as last 64 bits of the IPv6 address.",
		},
		"from_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the pool from which prefix will be taken to construct IPv6 address taking last part " +
				"of the address from address property.",
			AtLeastOneOf: []string{"address", "from_pool"},
		},
		"global": { // RO
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether address is global.",
		},
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"link_local": { //RO
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether address is link local.",
		},
		"no_dad": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If set indicates that address is anycast address and Duplicate Address Detection should " +
				"not be performed.",
		},
		"slave": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether address belongs to an interface which is a slave port to some other master interface",
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
