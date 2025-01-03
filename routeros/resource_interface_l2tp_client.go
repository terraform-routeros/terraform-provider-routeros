package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*3",
    "add-default-route": "false",
    "allow": "pap,chap,mschap1,mschap2",
    "allow-fast-path": "false",
    "connect-to": "1.2.3.4",
    "dial-on-demand": "false",
    "disabled": "false",
    "ipsec-secret": "",
    "keepalive-timeout": "60",
    "l2tp-proto-version": "l2tpv2",
    "l2tpv3-digest-hash": "md5",
    "max-mru": "1450",
    "max-mtu": "1450",
    "mrru": "disabled",
    "name": "l2tp-out1",
    "password": "bbb",
    "profile": "default-encryption",
    "running": "false",
    "use-ipsec": "false",
    "use-peer-dns": "no",
    "user": "aaa"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/2031631/L2TP#L2TP-L2TPClient
func ResourceInterfaceL2tpClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/l2tp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to add L2TP remote address as a default route.",
		},
		"allow": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Allowed authentication methods.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyAllowFastPath: PropAllowFastPathRw,
		"connect_to": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Remote address of L2TP server (if the address is in VRF table, VRF should be specified)" +
				" `/interface l2tp-client`\n`add connect-to=192.168.88.1@vrf1 name=l2tp-out1 user=l2tp-client`.",
		},
		KeyComment: PropCommentRw,
		"default_route_distance": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Since v6.2, sets distance value applied to auto created default route, if add-default-route " +
				"is also selected.",
			ValidateFunc: validation.IntBetween(1, 255),
		},
		"dial_on_demand": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Connects only when outbound traffic is generated. If selected, then route with gateway address " +
				"from `10.112.112.0/24` network will be added while connection is not established.",
		},
		KeyDisabled: PropDisabledRw,
		"ipsec_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Preshared key used when use-ipsec is enabled.",
		},
		"keepalive_timeout": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Since v6.0rc13, tunnel keepalive timeout in seconds.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"l2tpv3_circuit_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel.",
		},
		"l2tpv3_cookie_length": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Configures an L2TPv3 pseudowire static session cookie.",
			ValidateFunc:     validation.StringInSlice([]string{"0", "4-bytes", "8-bytes"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"l2tpv3_digest_hash": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which hash function to be used.",
			ValidateFunc:     validation.StringInSlice([]string{"md5", "none", "sha1"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"l2tp_proto_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specify protocol version.",
			ValidateFunc:     validation.StringInSlice([]string{"l2tpv2", "l2tpv3-ip", "l2tpv3-udp", "l2tpv"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mru": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Receive Unit. Max packet size that L2TP interface will be able to receive without " +
				"packet fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Transmission Unit. Max packet size that L2TP interface will be able to send without " +
				"packet fragmentation.",
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
		KeyName: PropName("Descriptive name of the interface."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used for authentication.",
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which PPP profile configuration will be used when establishing the tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify source address.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used for authentication.",
		},
		"use_ipsec": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When this option is enabled, dynamic IPSec peer configuration and policy (transport mode) " +
				"is added to encapsulate L2TP connection into IPSec tunnel. Multiple L2tp/ipsec clients behind the same " +
				"NAT will not work in this mode. To achieve such scenario, disable use-ipsec and set static policies " +
				"for clients with enabled `tunnel=yes`, `level=unique` settings.",
		},
		"use_peer_dns": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "To use peer dns.",
			ValidateFunc:     validation.StringInSlice([]string{"yes", "no", "exclusively"}, false),
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
