package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
iw-asra=any
iw-esr=any
iw-hessid=00:00:00:00:00:00
iw-hotspot20=any
iw-hotspot20-dgaf=any
iw-internet=any
iw-ipv4-availability=any
iw-ipv6-availability=any
iw-network-type=wildcard
iw-roaming-ois=""
iw-venue=any
iw-uesa=any
*/
// https://help.mikrotik.com/docs/spaces/ROS/pages/8978446/Wireless+Interface#WirelessInterface-ConnectList
// https://help.mikrotik.com/docs/spaces/ROS/pages/7962628/Interworking+Profiles
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
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_asra": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Additional Steps Required for Access. Set to yes, if a user should take additional steps " +
				"to access the internet, like the walled garden.",
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_authentication_types": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This property is only effective when `asra` is set to `yes`. Value of `url` is optional and " +
				"not needed if `dns-redirection` or `online-enrollment` is selected. To set the value of `url` to empty " +
				"string use double quotes. For example: `authentication-types=online-enrollment:\"\"`",
		},
		"iw_connection_capabilities": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This option allows to provide information about the allowed IP protocols and ports. This " +
				"information can be provided in ANQP response. The first number represents the IP protocol number, " +
				"the second number represents a port number.\n  * closed - set if protocol and port combination is not " +
				"allowed;\n  * open - set if protocol and port combination is allowed;\n  * unknown - set if protocol " +
				"and port combination is either open or closed.\n\nExample: `connection-capabilities=6:80:open,17:5060:closed`" +
				"Setting such a value on an Access Point informs the Wireless client, which is connecting to the " +
				"Access Point, that HTTP (6 - TCP, 80 - HTTP) is allowed and VoIP (17 - UDP; 5060 - VoIP) is not " +
				"allowed. This property does not restrict or allow usage of these protocols and ports, it only gives " +
				"information to station device which is connecting to Access Point.",
		},
		"iw_esr": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Emergency services reachable (ESR). Set to yes in order to indicate that emergency " +
				"services are reachable through the access point.",
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_hessid": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Homogenous extended service set identifier (HESSID). Devices that provide access to same " +
				"external networks are in one homogenous extended service set. This service set can be identified by " +
				"HESSID that is the same on all access points in this set. 6-byte value of HESSID is represented as MAC " +
				"address. It should be globally unique, therefore it is advised to use one of the MAC address of access " +
				"point in the service set.",
			ValidateFunc:     validation.IsMACAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_hotspot20": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Indicate Hotspot 2.0 capability of the Access Point.",
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_hotspot20_dgaf": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Downstream Group-Addressed Forwarding (DGAF). Sets value of DGAF bit to indicate whether " +
				"multicast and broadcast frames to clients are disabled or enabled.\n  * yes - multicast and broadcast " +
				"frames to clients are enabled;\n  * no - multicast and broadcast frames to clients are disabled.\n" +
				"To disable multicast and broadcast frames set `multicast-helper=full`.",
			ValidateFunc:     ValidationAutoYesNo,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_internet": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether the internet is available through this connection or not. This information is " +
				"included in the Interworking element.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_ipv4_availability": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Information about what IPv4 address and access are available." +
				"\n  * not-available - Address type not available;" +
				"\n  * public - public IPv4 address available;" +
				"\n  * port-restricted - port-restricted IPv4 address available;" +
				"\n  * single-nated - single NATed private IPv4 address available;" +
				"\n  * double-nated - double NATed private IPv4 address available;" +
				"\n  * port-restricted-single-nated -port-restricted IPv4 address and single NATed IPv4 address available;" +
				"\n  * port-restricted-double-nated - port-restricted IPv4 address and double NATed IPv4 address available;" +
				"\n  * unknown - availability of the address type is not known.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_ipv6_availability": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Information about what IPv6 address and access are available." +
				"\n  * not-available - Address type not available;" +
				"\n  * available - address type available;" +
				"\n  * unknown - availability of the address type is not known.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_network_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Information about network access type." +
				"\n  * emergency-only - a network dedicated and limited to accessing emergency services;" +
				"\n  * personal-device - a network of personal devices. An example of this type of network is a camera " +
				"that is attached to a printer, thereby forming a network for the purpose of printing pictures;" +
				"\n  * private - network for users with user accounts. Usually used in enterprises for employees, not guests;" +
				"\n  * private-with-guest - same as private, but guest accounts are available;" +
				"\n  * public-chargeable - a network that is available to anyone willing to pay. For example, a " +
				"subscription to Hotspot 2.0 service or in-room internet access in a hotel;" +
				"\n  * public-free - network is available to anyone without any fee. For example, municipal network in " +
				"city or airport Hotspot;" +
				"\n  * test - network used for testing and experimental uses. Not used in production;" +
				"\n  * wildcard - is used on Wireless clients. Sending probe request with a wildcard as network type " +
				"value will make all Interworking Access Points respond despite their actual network-type setting." +
				"\n\nA client sends a probe request frame with network-type set to value it is interested in. It will " +
				"receive replies only from access points with the same value (except the case of wildcard).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_realms": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Information about supported realms and the corresponding EAP method. " +
				"`realms=example.com:eap-tls,foo.ba:not-specified`",
		},
		"iw_roaming_ois": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Organization identifier (OI) usually are 24-bit is unique identifiers like organizationally " +
				"unique identifier (OUI) or company identifier (CID). In some cases, OI is longer for example OUI-36." +
				"A subscription service provider (SSP) can be specified by its OI. roaming-ois property can contain " +
				"zero or more SSPs OIs whose networks are accessible via this AP. Length of OI should be specified " +
				"before OI itself. For example, to set E4-8D-8C and 6C-3B-6B: `roaming-ois=03E48D8C036C3B6B`",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_uesa": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Unauthenticated emergency service accessible (UESA)." +
				"\n  * no - indicates that no unauthenticated emergency services are reachable through this Access Point;" +
				"\n  * yes - indicates that higher layer unauthenticated emergency services are reachable through this Access Point.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iw_venue": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify the venue in which the Access Point is located. Choose the value from available " +
				"ones. Some examples:\n   ```\n" +
				"  venue=business-bank\n" +
				"  venue=mercantile-shopping-mall\n" +
				"  venue=educational-university-or-college\n   ```",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
