package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "3gpp-info": "12",
    "authentication-types": "terms-and-conditions,terms-and-conditions",
    "disabled": "false",
    "domain-names": "example.com,www.example.com",
    "esr": "true",
    "hessid": "00:00:00:00:00:00",
    "hotspot20": "true",
    "hotspot20-dgaf": "true",
    "internet": "true",
    "ipv4-availability": "not-available",
    "ipv6-availability": "unknown",
    "name": "interworking1",
    "network-type": "private",
    "operational-classes": "10,11",
    "operator-names": "name1,name2",
    "realms": "something1:not-specified,something2:eap-sim",
    "roaming-ois": "something1,something2",
    "uesa": "true",
    "venue": "unspecified",
    "venue-names": "name1,name2",
    "wan-at-capacity": "true",
    "wan-downlink": "10",
    "wan-downlink-load": "5",
    "wan-measurement-duration": "10",
    "wan-status": "reserved",
    "wan-symmetric": "true",
    "wan-uplink": "2",
    "wan-uplink-load": "1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Interworking+Profiles
func ResourceWifiInterworking() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/interworking"),
		MetaId:           PropId(Id),

		"3gpp_info": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Cellular network advertisement information - country and network codes.",
		},
		"3gpp_raw": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cellular network advertisement information - country and network codes.",
		},
		"asra": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable Additional Steps Required for Access.",
		},
		"authentication_types": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of authentication types that is only effective when `asra` is set to yes.",
		},
		KeyComment: PropCommentRw,
		"connection_capabilities": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list to provide information about the allowed IP protocols and ports.",
		},
		KeyDisabled: PropDisabledRw,
		"domain_names": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of fully qualified domain names (FQDN) that indicate the entity operating the Hotspot.",
		},
		"esr": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable Emergency Services Reachability.",
		},
		"hessid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Homogenous extended service set identifier (HESSID).",
		},
		"hotspot20": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to indicate Hotspot 2.0 capability of the Access Point.",
		},
		"hotspot20_dgaf": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to indicate Downstream Group-Addressed Forwarding (DGAF) capability.",
		},
		"internet": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to indicate Internet availability.",
		},
		"ipv4_availability": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to indicate IPv4 availability.",
			ValidateFunc: validation.StringInSlice([]string{"double-nated", "not-available", "port-restricted", "port-restricted-double-nated", "port-restricted-single-nated", "public", "single-nated", "unknown"}, false),
		},
		"ipv6_availability": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to indicate IPv6 availability.",
			ValidateFunc: validation.StringInSlice([]string{"available", "not-available", "unknown"}, false),
		},
		KeyName: PropName("Name of the interworking profile."),
		"network_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Information about network access type.",
			ValidateFunc: validation.StringInSlice([]string{"emergency-only", "personal-device", "private", "private-with-guest", "public-chargeable", "public-free", "test", "wildcard"}, false),
		},
		"operational_classes": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
			Description: "A list with information about other available bands.",
		},
		"operator_names": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of colon-separated operator names and language codes.",
		},
		"realms": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of colon-separated realm names and EAP methods.",
		},
		"realms_raw": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of 'NAI Realm Tuple' excluding 'NAI Realm Data Field Length' field.",
		},
		"roaming_ois": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of Organization Identifiers (OI).",
		},
		"uesa": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable Unauthenticated Emergency Service Accessibility.",
		},
		"venue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Information about the venue in which the Access Point is located.",
		},
		"venue_names": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "A list of colon-separated venue names and language codes.",
		},
		"wan_at_capacity": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to indicate that the Access Point or the network is at its max capacity.",
		},
		"wan_downlink": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The downlink speed of the WAN connection set in kbps.",
		},
		"wan_downlink_load": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The downlink load of the WAN connection measured over `wan_measurement_duration`.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"wan_measurement_duration": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The duration during which `wan_downlink_load` and `wan_uplink_load` are measured.",
			ValidateFunc: Validation64k,
		},
		"wan_status": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Information about the status of the Access Point's WAN connection.",
			ValidateFunc: validation.StringInSlice([]string{"down", "reserved", "test", "up"}, false),
		},
		"wan_symmetric": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to indicate that the WAN link is symmetric (upload and download speeds are the same).",
		},
		"wan_uplink": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The uplink speed of the WAN connection set in kbps.",
		},
		"wan_uplink_load": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The uplink load of the WAN connection measured over `wan_measurement_duration`.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
