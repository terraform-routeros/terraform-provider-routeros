package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://help.mikrotik.com/docs/spaces/ROS/pages/8978446/Wireless+Interface#WirelessInterface-ConnectList
func ResourceInterfaceWirelessConnectList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless/connect-list"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("three_gpp:3gpp"),

		"three_gpp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
		"allow_signal_out_of_range": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: TimeEquall,
		},
		"area_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rule matches if area value of AP (a proprietary extension) begins with specified value.area " +
				"value is a proprietary extension.",
		},
		KeyComment: PropCommentRw,
		"connect": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Available options: yes - Connect to access point that matches this rule. no - Do not connect " +
				"to any access point that matches this rule.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled:  PropDisabledRw,
		KeyInterface: PropInterfaceRw,
		"interworking": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			ValidateFunc:     validation.StringInSlice([]string{"yes", "no", "any"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		/* ???
		        iw-asra                        iw-esr           iw-hotspot20-dgaf        iw-ipv6-availability     iw-roaming-ois
				iw-authentication-types        iw-hessid        iw-internet              iw-network-type          iw-uesa
				iw-connection-capabilities     iw-hotspot20     iw-ipv4-availability     iw-realms                iw-venue
		*/
		KeyMacAddress: PropMacAddressRw("Rule matches only AP with the specified MAC address.", false),
		"security_profile": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of security profile that is used when connecting to matching access points, If value " +
				"of this property is none, then security profile specified in the interface configuration will be used. " +
				"In station mode, rule will match only access points that can support specified security profile. Value " +
				"none will match access point that supports security profile that is specified in the interface configuration. " +
				"In access point mode value of this property will not be used to match remote devices.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"signal_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rule matches if signal strength of the access point is within the range. If station establishes " +
				"connection to access point that is matched by this rule, it will disconnect from that access point when " +
				"signal strength goes out of the specified range.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssid": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rule matches access points that have this SSID. Empty value matches any SSID. This property " +
				"has effect only when station mode interface ssid is empty, or when access point mode interface has wds-ignore-ssid=yes.",
		},
		"wireless_protocol": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			ValidateFunc:     validation.StringInSlice([]string{"802.11", "any", "nstreme", "tdma"}, false),
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
