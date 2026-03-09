package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "enabled": "false",
  "authentication": "mschap1,mschap2",
  "default-profile": "default-encryption",
  "mrru": "disabled",
	"max-mru": "1450",
  "keepalive-timeout": "30",
  "max-mtu": "1450",
  "use-ipsec": "false",
	"ipsec-secret": "",
	"accept-proto-version", "all",
  "accept-pseudowire-type": "all",
	"allow-fast-path": "false",
	"caller-id-type": "ip-address",
	"max-sessions": "unlimited",
	"one-session-per-host": "false",
	"l2tpv3-circuit-id": "",
	"l2tpv3-cookie-length", "0",
	"l2tpv3-digest-hash", "md5",
	"l2tpv3-ether-interface-list": ""
}
*/

// https://help.mikrotik.com/docs/display/ROS/L2TP#L2TP-L2TPServer
func ResourceInterfaceL2tpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/l2tp-server/server"),
		MetaId:           PropId(Id),

    KeyEnabled: PropEnabled("Enables/disables service."),
		"authentication": {
			Type:             schema.TypeSet,
			Optional:         true,
			Description:      "Authentication algorithm.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				Default:      "all",
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
		},
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Default profile to use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, " +
				"it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the " +
				"tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mru": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Receive Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(512, 18432),
		},
		"keepalive_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "30",
			Description:      "Sets keepalive timeout in seconds.",
			DiffSuppressFunc: TimeEqual,
		},
		"max_mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Transmission Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(512, 18432),
		},
		"use_ipsec": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When this option is enabled, dynamic IPSec peer configuration is added to suite" +
			  "most of the L2TP road-warrior setups. When require is selected server will accept only" +
				"those L2TP connection attempts that were encapsulated in the IPSec tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipsec_secret": {
			Type: schema.TypeString,
			Optional: true,
			Description: "Preshared key used when use-ipsec is enabled.",
		  DiffSuppressFunc: AlwaysPresentNotUserProvided,	
		},
		"accept_proto_version" : {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specify protocol version.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"all", "l2tpv2", "l2tpv3"}, false),
		},
    "accept_pseudowire_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Set the pseudowire signaling protocol for specific pseudowire type.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"all", "ether", "ppp"}, false),
		},
    "allow_fast_path": {
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Description:      "To forward packets without additional processing in the Linux kernel.",
		},
		"caller_id_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "If same source IP address is used for multiple clients set id type to number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"ip-address", "number"}, false),
		},
		"max_sessions": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Set number of needed sessions.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"one_session_per_host": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "To allow one session per host.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"l2tpv3_circuit_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"l2tpv3_cookie_length": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Configures an L2TP pseudowire static session cookie.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"0", "4-bytes", "8-bytes"}, false),
		},
    "l2tpv3_digest_hash": {
		  Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which hash function to be used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"md5", "none", "sha1"}, false),	
		},
    "l2tpv3_ether_interface_list": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Set your interface list for example the default ones- all, dynamic, none, static.",
		  DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
