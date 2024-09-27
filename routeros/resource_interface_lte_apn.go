package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "add-default-route": "true",
    "apn": "internet",
    "authentication": "none",
    "comment": "wan",
    "default": "true",
    "default-route-distance": "2",
    "ip-type": "auto",
    "name": "default",
    "passthrough-interface": "ether2",
    "passthrough-mac": "00:00:00:00:00:00",
    "use-network-apn": "true",
    "use-peer-dns": "false"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/LTE#LTE-APNprofiles
func ResourceInterfaceLteApn() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/lte/apn"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to add a default route to forward all traffic over the LTE interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"apn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Service Provider's Access Point Name.",
		},
		"authentication": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allowed protocol to use for authentication.",
			ValidateFunc:     validation.StringInSlice([]string{"pap", "chap", "none"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		KeyDefault: PropDefaultRo,
		"default_route_distance": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Sets distance value applied to auto-created default route, if add-default-route is also " +
				"selected. LTE route by default is with distance 2 to prefer wired routes over LTE.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ip_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Requested PDN type.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "ipv4", "ipv4-ipv6", "ipv6"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipv6_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface on which to advertise IPv6 prefix.",
		},
		KeyName: PropName("APN profile name"),
		"number": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "APN profile number.",
		},
		"passthrough_interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interface to passthrough IP configuration (activates passthrough).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"passthrough_mac": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "If set to auto, then will learn MAC from the first packet.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"passthrough_subnet_selection": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "`auto` selects the smallest possible subnet to be used for the passthrough interface. `p2p` " +
				"sets the passthrough interface subnet as `/32` and picks gateway address from `10.177.0.0/16` range. " +
				"The gateway address stays the same until the apn configuration is changed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Password used if any of the authentication protocols are active.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_network_apn": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Parameter is available starting from RouterOS v7 and used only for MBIM modems. If set to yes, " +
				"uses network provided APN.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_peer_dns": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "If set to yes, uses DNS received from LTE interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Username used if any of the authentication protocols are active.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
