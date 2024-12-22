package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*AD",
    "comment": "comment",
    "disabled": "false",
    "name": "pppoe-in1",
    "running": "false",
    "service": "",
    "user": ""
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/2031625/PPPoE#PPPoE-PPPoEServer
func ResourceInterfacePppoeServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pppoe-server"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Authentication algorithm.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyInterface: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface that the clients are connected to",
		},
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Defines the time period (in seconds) after which the router is starting to send keepalive " +
				"packets every second. If there is no traffic and no keepalive responses arrive for that period of time " +
				"(i.e. 2 * keepalive-timeout), the non responding client is proclaimed disconnected.",
			DiffSuppressFunc: TimeEquall,
		},
		"max_mru": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Receive Unit. The optimal value is the MTU of the interface the tunnel is working " +
				"over reduced by 20 (so, for 1500-byte Ethernet link, set the MTU to 1480 to avoid fragmentation of packets).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Transmission Unit. The optimal value is the MTU of the interface the tunnel is working " +
				"over reduced by 20 (so, for 1500-byte Ethernet link, set the MTU to 1480 to avoid fragmentation of packets).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_sessions": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of clients that the AC can serve. '0' = no limitations.",
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
		KeyName: PropName(""),
		"one_session_per_host": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow only one session per host (determined by MAC address). If a host tries to establish " +
				"a new session, the old one will be closed.",
		},
		"pppoe_over_vlan_range": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "This setting allows a PPPoE server to operate over 802.1Q VLANs. By default, a PPPoE server " +
				"only accepts untagged packets on its interface. However, in scenarios where clients are on separate " +
				"VLANs, instead of creating multiple 802.1Q VLAN interfaces and bridging them together or configuring " +
				"individual PPPoE servers for each VLAN, you can specify the necessary VLANs directly in the PPPoE server " +
				"settings. When you specify the VLAN IDs, the PPPoE server will accept both untagged packets and 802.1Q " +
				"tagged packets from clients, and it will reply using the same VLAN. This setting can also be applied " +
				"to both CVLAN and SVLAN interfaces. For example, when the use-service-tag=yes option is used on a VLAN " +
				"interface, enabling QinQ setups as well. The setting supports a range of VLAN IDs, as well as individual " +
				"VLANs specified using comma-separated values. For example: pppoe-over-vlan-range=100-115,120,122,128-130.",
		},
		KeyRunning: PropRunningRo,
		"service": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "This attribute is required in the ROS 7 version.",
		},
		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The PPPoE service name. Server will accept clients which sends PADI message with service-names " +
				"that matches this setting or if service-name field in PADI message is not set.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "This attribute is required in the ROS 7 version.",
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
