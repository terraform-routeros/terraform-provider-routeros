package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "dynamic": "false",
    "exchange-mode": "main",
    "name": "peer1",
    "passive": "true",
    "profile": "default",
    "responder": "true",
    "send-initial-contact": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Peers
func ResourceIpIpsecPeer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/peer"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If the remote peer's address matches this prefix, then the peer configuration is used in authentication " +
				"and establishment of Phase 1. If several peer's addresses match several configuration entries, the most " +
				"specific one (i.e. the one with the largest netmask) will be used.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"exchange_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Different ISAKMP phase 1 exchange modes according to RFC 2408. the main mode relaxes rfc2409 " +
				"section 5.4, to allow pre-shared-key authentication in the main mode. ike2 mode enables Ikev2 RFC 7296. " +
				"Parameters that are ignored by IKEv2 proposal-check, compatibility-options, lifebytes, dpd-maximum-failures, " +
				"nat-traversal.",
			ValidateFunc:     validation.StringInSlice([]string{"aggressive", "base", "main", "ike2"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"local_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Routers local address on which Phase 1 should be bounded to.",
		},
		KeyName: PropName("Peer name."),
		"passive": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When a passive mode is enabled will wait for a remote peer to initiate an IKE connection. " +
				"The enabled passive mode also indicates that the peer is xauth responder, and disabled passive mode " +
				"- xauth initiator. When a passive mode is a disabled peer will try to establish not only phase1 but " +
				"also phase2 automatically, if policies are configured or created during the phase1.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Communication port used (when a router is an initiator) to connect to remote peer in cases " +
				"if remote peer uses the non-default port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the profile template that will be used during IKE negotiation.",
			ValidateFunc:     validation.StringInSlice([]string{"string"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"responder": {
			Type:     schema.TypeBool,
			Computed: true,
			Description: "Whether this peer will act as a responder only (listen to incoming requests) and not " +
				"initiate a connection.",
		},
		"send_initial_contact": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies whether to send `initial contact` IKE packet or wait for remote side, this packet " +
				"should trigger the removal of old peer SAs for current source address. Usually, in road warrior setups " +
				"clients are initiators and this parameter should be set to no. Initial contact is not sent if modecfg " +
				"or xauth is enabled for ikev1.",
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
