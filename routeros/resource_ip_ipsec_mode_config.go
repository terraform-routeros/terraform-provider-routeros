package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*3",
    "address-pool": "default-dhcp",
    "address-prefix-length": "24",
    "name": "cfg1",
    "responder": "true",
    "split-dns": "1.1.1.1",
    "split-include": "0.0.0.0/0",
    "system-dns": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Modeconfigs
func ResourceIpIpsecModeConfig() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/mode-config"),
		MetaId:           PropId(Id),

		"address": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "Single IP address for the initiator instead of specifying a whole address pool.",
			ValidateFunc:  validation.IsIPv4Address,
			ConflictsWith: []string{"address_pool"},
		},
		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the address pool from which the responder will try to assign address if mode-config " +
				"is enabled.",
			ConflictsWith: []string{"address"},
		},
		"address_prefix_length": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Prefix length (netmask) of the assigned address from the pool.",
			ValidateFunc:     validation.IntBetween(1, 32),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"connection_mark": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Firewall connection mark.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// KeyComment: PropCommentRw,
		KeyName: PropName(""),
		"responder": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether the configuration will work as an initiator (client) or responder (server). " +
				"The initiator will request for mode-config parameters from the responder.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"split_dns": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "List of DNS names that will be resolved using a system-dns=yes or static-dns= setting.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
		},
		"split_include": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "List of subnets in CIDR format, which to tunnel. Subnets will be sent to the peer using the " +
				"CISCO UNITY extension, a remote peer will create specific dynamic policies.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		"src_address_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifying an address list will generate dynamic source NAT rules. This parameter is only " +
				"available with responder=no. A roadWarrior client with NAT.",
		},
		"static_dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Manually specified DNS server's IP address to be sent to the client.",
		},
		"system_dns": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "When this option is enabled DNS addresses will be taken from `/ip dns`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_responder_dns": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			ValidateFunc:     validation.StringInSlice([]string{"exclusively", "yes", "no"}, false),
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
