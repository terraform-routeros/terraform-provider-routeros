package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
 {
  ".id": "*1C",
  "ac-name": "",
  "add-default-route": "true",
  "allow": "pap,chap,mschap1,mschap2",
  "default-route-distance": "1",
  "dial-on-demand": "false",
  "disabled": "false",
  "interface": "ether2",
  "invalid": "false",
  "keepalive-timeout": "10",
  "max-mru": "auto",
  "max-mtu": "auto",
  "mrru": "disabled",
  "name": "pppoe-out1",
  "password": "",
  "profile": "default",
  "running": "false",
  "service-name": "",
  "use-peer-dns": "false",
  "user": ""
 }
*/

// https://help.mikrotik.com/docs/display/ROS/PPPoE#PPPoE-PPPoEClient
func ResourceInterfacePPPoEClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pppoe-client"),
		MetaId:           PropId(Id),

		"ac_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
			Description: "Access Concentrator name, this may be left blank and the client will connect to any " +
				"access concentrator on the broadcast domain.",
		},
		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enable/Disable whether to add default route automatically.",
		},
		"allow": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Description: "Allowed authentication methods, by default all methods are allowed.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
		},
		KeyComment: PropCommentRw,
		"default_route_distance": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
			Description: "sets distance value applied to auto created default route, if add-default-route is also " +
				"selected.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"dial_on_demand": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "connects to AC only when outbound traffic is generated. If selected, then route with " +
				"gateway address from 10.112.112.0/24 network will be added while connection is not " +
				"established.",
		},
		KeyDisabled: PropDisabledRw,
		"host_uniq": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: HexEqual,
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface name on which client will run.",
		},
		KeyInvalid: PropInvalidRo,
		"keepalive_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     10,
			Description: "Sets keepalive timeout in seconds.",
		},
		"max_mru": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "auto",
			Description: "Maximum Receive Unit.",
		},
		"max_mtu": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "auto",
			Description: "Maximum Transmission Unit.",
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "disabled",
			Description: "Maximum packet size (512..65535 or disabled) that can be received on the link. If a packet " +
				"is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet " +
				"packets to be sent over the tunnel.",
		},
		KeyName: PropName("Name of the PPPoE interface."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: "Password used to authenticate.",
		},
		"profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			Description: "Specifies which PPP profile configuration will be used when establishing the tunnel.",
		},
		KeyRunning: PropRunningRo,
		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
			Description: "Specifies the service name set on the access concentrator, can be left blank to connect " +
				"to any PPPoE server.",
		},
		"use_peer_dns": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enable/disable getting DNS settings from the peer.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Username used for authentication.",
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
