package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServerOptionMatcher https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP#DHCP-Optionmatcher
func ResourceDhcpServerOptionMatcher() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/matcher"),
		MetaId:           PropId(Id),

		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IP pool, from which to take IP addresses for the clients. " +
				"If set to static-only, then only the clients that have a static " +
				"lease (added in lease submenu) will be allowed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"code": {
			Type:     schema.TypeInt,
			Required: true,
			Description: "DHCP option code. All codes are available at " +
				"http://www.iana.org/assignments/bootp-dhcp-parameters",
			ValidateFunc: validation.IntBetween(1, 254),
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"matching_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matching method:\n\n" +
				"- exact: option should match exactly to value\n" +
				"- substring: value can match anywhere in the option string; " +
				"at the start, middle, or end.",
			ValidateFunc: validation.StringInSlice(
				[]string{"exact", "substring"},
				false,
			),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropNameForceNewRw,
		"option_set": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "A custom set of DHCP options defined in the Option Sets menu.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Server name which serves option matcher.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"value": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A value that will be searched for in option.\n" +
				"Available data types for value are:\n\n" +
				"- string\n" +
				"- HEX",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
