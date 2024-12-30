package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceDhcpRelay https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Relay
func ResourceDhcpRelay() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-relay"),
		MetaId:           PropId(Id),

		"add_relay_info": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Adds DHCP relay agent information if enabled according to RFC 3046. Agent Circuit ID " +
				"Sub-option contains mac address of an interface, Agent Remote ID Sub-option contains MAC address " +
				"of the client from which request was received.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"delay_threshold": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyDisabled: PropDisabledRw,
		"dhcp_server": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "List of DHCP servers' IP addresses which should the DHCP requests be forwarded to.",
		},
		"dhcp_server_vrf": PropVrfRw,
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface name the DHCP relay will be working on.",
		},
		KeyInvalid: PropInvalidRo,
		"local_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The unique IP address of this DHCP relay needed for DHCP server to distinguish relays. " +
				"If set to 0.0.0.0 - the IP address will be chosen automatically",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Descriptive name for the relay."),
		"relay_info_remote_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specified string will be used to construct Option 82 instead of client's MAC address. Option " +
				"82 consist of: interface from which packets was received + client mac address or relay-info-remote-id",
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
