package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceInterfaceWireguardPeer https://help.mikrotik.com/docs/display/ROS/WireGuard#WireGuard-Peers
func ResourceInterfaceWireguardPeer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireguard/peers"),
		MetaId:           PropId(Id),

		"allowed_address": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer " +
				"is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be " +
				"specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: ValidationIpAddress,
			},
		},
		"current_endpoint_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The most recent source IP address of correctly authenticated packets from the peer.",
		},
		"current_endpoint_port": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The most recent source IP port of correctly authenticated packets from the peer.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"endpoint_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An endpoint IP or hostname can be left blank to allow remote connection from any address.",
		},
		"endpoint_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An endpoint port can be left blank to allow remote connection from any port.",
		},
		KeyInterface: PropInterfaceRw,
		"last_handshake": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time in seconds after the last successful handshake.",
		},
		"persistent_keepalive": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated " +
				"empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid " +
				"persistently. For example, if the interface very rarely sends traffic, but it might at anytime " +
				"receive traffic from a peer, and it is behind NAT, the interface might benefit from having a " +
				"persistent keepalive interval of 25 seconds.",
		},
		"preshared_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer " +
				"of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for " +
				"post-quantum resistance.",
		},

		"public_key": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The remote peer's calculated public key.",
		},
		"rx": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The total amount of bytes received from the peer.",
		},
		"tx": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The total amount of bytes transmitted to the peer.",
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
