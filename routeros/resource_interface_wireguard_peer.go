package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceWireguardPeer https://help.mikrotik.com/docs/display/ROS/WireGuard#WireGuard-Peers
func ResourceInterfaceWireguardPeer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireguard/peers"),
		MetaId:           PropId(Id),

		"allowed_address": {
			Type:     schema.TypeList,
			Required: true,
			Description: "List of IP (v4 or v6) addresses with CIDR masks from which incoming traffic for this peer " +
				"is allowed and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be " +
				"specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6 addresses.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				// ValidateFunc: ValidationIpAddress,
			},
		},
		"client_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When imported using a qr code for a client (for example, a phone), then this address for the " +
				"wg interface is set on that device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"client_dns": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specify when using WireGuard Server as a VPN gateway for peer traffic.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"client_endpoint": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The IP address and port number of the WireGuard Server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"client_keepalive": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Same as persistent-keepalive but from peer side.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"client_listen_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The local port upon which this WireGuard tunnel will listen for incoming traffic from peers, " +
				"and the port from which it will source outgoing packets.",
			ValidateFunc:     validation.IntBetween(0, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
		KeyDynamic:  PropDynamicRo,
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
		KeyName: PropNameOptional("Name of the tunnel."),
		"persistent_keepalive": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated " +
				"empty packet to the peer for the purpose of keeping a stateful firewall or NAT mapping valid " +
				"persistently. For example, if the interface very rarely sends traffic, but it might at anytime " +
				"receive traffic from a peer, and it is behind NAT, the interface might benefit from having a " +
				"persistent keepalive interval of 25 seconds.",
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^\d+s$`),
				"value should be an integer between 1 and 65535 inclusive: 5s, 25s, ...",
			),
		},
		"preshared_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "A **base64** preshared key. Optional, and may be omitted. This option adds an additional layer " +
				"of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for " +
				"post-quantum resistance.",
		},
		"private_key": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "A base64 private key. If not specified, it will be automatically generated upon interface creation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
