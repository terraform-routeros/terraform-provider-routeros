package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1000001",
    "action": "encrypt",
    "active": "false",
    "disabled": "false",
    "dst-address": "::/0",
    "dst-port": "any",
    "dynamic": "false",
    "invalid": "false",
    "ipsec-protocols": "esp",
    "level": "require",
    "peer": "peer1",
    "ph2-count": "0",
    "ph2-state": "no-phase2",
    "proposal": "default",
    "protocol": "all",
    "sa-dst-address": "::",
    "sa-src-address": "::",
    "src-address": "::/0",
    "src-port": "any",
    "tunnel": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Policies
func ResourceIpIpsecPolicy() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/policy"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("ph2_count", "ph2_state", "sa_dst_address", "sa_src_address"),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies what to do with the packet matched by the policy.none - pass the packet unchanged.discard " +
				"- drop the packet.encrypt - apply transformations specified in this policy and it's SA.",
			ValidateFunc:     validation.StringInSlice([]string{"discard", "encrypt", "none"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Destination address to be matched in packets. Applicable when tunnel mode (`tunnel=yes`) or " +
				"template (`template=yes`) is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Destination port to be matched in packets. If set to any all ports will be matched.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDynamic: PropDynamicRo,
		"group": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the policy group to which this **template** is assigned.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInvalid: PropInvalidRo,
		"ipsec_protocols": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies what combination of Authentication Header and Encapsulating Security Payload protocols " +
				"you want to apply to matched traffic.",
			ValidateFunc:     validation.StringInSlice([]string{"ah", "esp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"level": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies what to do if some of the SAs for this policy cannot be found:\n  * use - skip this transform, " +
				"do not drop the packet, and do not acquire SA from IKE daemon;\n  * require - drop the packet and acquire " +
				"SA;\n  * unique - drop the packet and acquire a unique SA that is only used with this particular policy. It " +
				"is used in setups where multiple clients can sit behind one public IP address (clients behind NAT).",
			ValidateFunc:     validation.StringInSlice([]string{"require", "unique", "use"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"peer": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the peer on which the policy applies.",
		},
		"proposal": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the proposal template that will be sent by IKE daemon to establish SAs for this policy.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP packet protocol to match.",
			ValidateFunc: validation.StringInSlice([]string{"all", "dccp", "ddp", "egp", "encap", "etherip", "ggp",
				"gre", "hmp", "icmp", "icmpv6", "idpr-cmtp", "igmp", "ipencap", "ipip", "ipsec-ah", "ipsec-esp",
				"ipv6-encap", "ipv6-frag", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "iso-tp4", "l2tp", "ospf", "pim",
				"pup", "rdp", "rspf", "rsvp", "sctp", "st", "tcp", "udp", "udp-lite", "vmtp", "vrrp", "xns-idp", "xtp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Source address to be matched in packets. Applicable when tunnel mode (`tunnel=yes`) or template " +
				"(`template=yes`) is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Source port to be matched in packets. If set to any all ports will be matched.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"template": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Creates a template and assigns it to a specified policy group.Following parameters are used " +
				"by template:\n  * group - name of the policy group to which this template is assigned;\n  * src-address,\n  * dst-address " +
				"- Requested subnet must match in both directions (for example 0.0.0.0/0 to allow all);\n  * protocol - protocol " +
				"to match, if set to all, then any protocol is accepted;\n  * proposal - SA parameters used for this template;\n  * level " +
				"- useful when unique is required in setups with multiple clients behind NAT.",
		},
		"tunnel": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Specifies whether to use tunnel mode.",
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
