package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "caller-id": "",
    "disabled": "false",
    >>>"ipv6-routes": "",
    "last-caller-id": "172.18.0.2",
    "last-disconnect-reason": "hung-up",
    "last-logged-out": "may/01/2023 20:52:13",
    "limit-bytes-in": "0",
    "limit-bytes-out": "0",
    "local-address": "172.18.0.2",
    "name": "user1",
    "password": "1",
    "profile": "ovpn",
    "routes": "",
    "service": "ovpn"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/PPP+AAA#PPPAAA-UserDatabase
func ResourcePPPSecret() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ppp/secret"),
		MetaId:           PropId(Id),

		"caller_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "For PPTP and L2TP it is the IP address a client must connect from. For PPPoE it is the " +
				"MAC address (written in CAPITAL letters) a client must  connect from. For ISDN it is the " +
				"caller's number (that may or may not be  provided by the operator) the client may " +
				"dial-in from.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"ipv6_routes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "IPv6 routes.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"last_caller_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_disconnect_reason": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_logged_out": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"limit_bytes_in": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Maximal amount of bytes for a session that client can upload.",
		},
		"limit_bytes_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Maximal amount of bytes for a session that client can download.",
		},
		"local_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "IP address that will be set locally on ppp interface.",
			ValidateFunc: validation.IsIPv4Address,
		},
		KeyName: PropName("Name used for authentication."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Password used for authentication.",
			Sensitive:   true,
		},
		"profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			Description: "Which user profile to use.",
		},
		"remote_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "IP address that will be assigned to remote ppp interface.",
			ValidateFunc: validation.IsIPv4Address,
		},
		"remote_ipv6_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IPv6 prefix assigned to ppp client. Prefix is added to ND prefix list enabling stateless " +
				"address auto-configuration on ppp interface.Available starting from v5.0.",
		},
		"routes": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Routes  that appear on the server when the client is connected. The route  format is: " +
				"dst-address gateway metric (for example, 10.1.0.0/ 24  10.0.0.1 1). Other syntax is not " +
				"acceptable since it can be represented  in incorrect way. Several routes may be " +
				"specified separated with commas.  This parameter will be ignored for OpenVPN.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		// The ROS 7.8 version does not contain the isdn option.
		"service": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "any",
			Description: "Specifies the services that particular user will be able to use.",
			ValidateFunc: validation.StringInSlice(
				[]string{"any", "async", "isdn", "l2tp", "pppoe", "pptp", "ovpn", "sstp"},
				false,
			),
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
