package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://help.mikrotik.com/docs/display/ROS/Dot1X#Dot1X-Client
func ResourceInterfaceDot1xClientV0() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/dot1x/client"),
		MetaId:           PropId(Id),

		"anon_identity": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Identity for outer layer EAP authentication. Used only with `eap-ttls` and `eap-peap` methods. If not set, the value from the identity parameter will be used for outer layer EAP authentication.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Name of a certificate. Required when the `eap-tls` method is used.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"eap_methods": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A list of EAP methods used for authentication: `eap-tls`, `eap-ttls`, `eap-peap`, `eap-mschapv2`.",
		},
		"identity": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The supplicant identity that is used for EAP authentication.",
		},
		KeyInterface: PropInterfaceRw,
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Cleartext password for the supplicant.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
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

// https://help.mikrotik.com/docs/display/ROS/Dot1X#Dot1X-Server
func ResourceInterfaceDot1xServerV0() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/dot1x/server"),
		MetaId:           PropId(Id),

		"accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to send RADIUS accounting requests to the authentication server.",
		},
		"auth_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "1m",
			Description:      "Total time available for EAP authentication.",
			DiffSuppressFunc: TimeEquall,
		},
		"auth_types": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "dot1x",
			Description: "Used authentication type on a server interface. Comma-separated list of `dot1x` and `mac-auth`.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"guest_vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Assigned VLAN when end devices do not support dot1x authentication and no mac-auth fallback is configured.",
			ValidateFunc: validation.IntBetween(1, 4094),
		},
		KeyInterface: PropInterfaceRw,
		"interim_update": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "Interval between scheduled RADIUS Interim-Update messages.",
			DiffSuppressFunc: TimeEquall,
		},
		"mac_auth_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "mac-as-username",
			Description:  "An option that allows to control User-Name and User-Password RADIUS attributes when using MAC authentication.",
			ValidateFunc: validation.StringInSlice([]string{"mac-as-username", "mac-as-username-and-password"}, false),
		},
		"radius_mac_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "XX:XX:XX:XX:XX:XX",
			Description: "An option that controls how the MAC address of the client is encoded in the User-Name and User-Password attributes when using MAC authentication.",
			ValidateFunc: validation.StringInSlice([]string{"XX-XX-XX-XX-XX-XX", "XX:XX:XX:XX:XX:XX", "XXXXXXXXXXXX",
				"xx-xx-xx-xx-xx-xx", "xx:xx:xx:xx:xx:xx", "xxxxxxxxxxxx"}, false),
		},
		"reauth_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option that enables server port re-authentication.",
			DiffSuppressFunc: TimeEquall,
		},
		"reject_vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Assigned VLAN when authentication failed, and a RADIUS server responded with an Access-Reject message. ",
			ValidateFunc: validation.IntBetween(1, 4094),
		},
		"retrans_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "30s",
			Description:      "The time interval between message re-transmissions if no response is received from the supplicant.",
			DiffSuppressFunc: TimeEquall,
		},
		"server_fail_vlan_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Assigned VLAN when RADIUS server is not responding and request timed out.",
			ValidateFunc: validation.IntBetween(1, 4094),
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
