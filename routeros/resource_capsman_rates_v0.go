package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceCapsManRatesV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/caps-man/rates"),
			MetaId:           PropId(Name),

			"basic": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "List of basic rates. Client will connect to AP only if it supports all basic " +
					"rates announced by the AP. AP will establish WDS link only if it supports all basic " +
					"rates of the other AP.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"1Mbps", "2Mbps", "5.5Mbps", "6Mbps",
						"9Mbps", "11Mbps", "12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps", "54Mbps"}, false),
				},
			},
			"supported": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "List of supported rates. Two devices will communicate only using rates that " +
					"are supported by both devices.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"1Mbps", "2Mbps", "5.5Mbps", "6Mbps",
						"9Mbps", "11Mbps", "12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps", "54Mbps"}, false),
				},
			},
			KeyComment: PropCommentRw,
			"ht_basic_mcs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Description: "Modulation and Coding Schemes that every connecting client must support. Refer to " +
					"802.11n for MCS specification.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringMatch(regexp.MustCompile(`mcs-\d+`),
						`ht_basic_mcs format is "mcs-[0..23]": mcs-"12"`),
				},
			},
			"ht_supported_mcs": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n " +
					"for MCS specification.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringMatch(regexp.MustCompile(`mcs-\d+`),
						`ht_supported_mcs format is "mcs-[0..23]": "mcs-11"`),
				},
			},
			KeyName: PropNameForceNewRw,
			"vht_basic_mcs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Modulation and Coding Schemes that every connecting client must support. Refer to " +
					"802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none " +
					"- will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS " +
					"0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9",
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"none", "mcs0-7", "mcs0-8", "mcs0-9"}, false, false),
			},
			"vht_supported_mcs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to " +
					"802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none " +
					"- will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported " +
					"MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - " +
					"devices will advertise as supported MCS-0 to MCS-9",
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"none", "mcs0-7", "mcs0-8", "mcs0-9"}, false, false),
			},
		},
	}
}
