package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "bridge-type": "customer-vid-used-as-lookup-vid",
  "bypass-ingress-port-policing-for": "",
  "bypass-l2-security-check-filter-for": "",
  "bypass-vlan-ingress-filter-for": "",
  "drop-if-invalid-or-src-port-not-member-of-vlan-on-ports": "ether1,ether2,ether3,ether4,ether5,ether6,ether7,ether8",
  "drop-if-no-vlan-assignment-on-ports": "",
  "egress-mirror-ratio": "1/1",
  "egress-mirror0": "switch1-cpu,modified",
  "egress-mirror1": "switch1-cpu,modified",
  "fdb-uses": "mirror0",
  "forward-unknown-vlan": "true",
  "ingress-mirror-ratio": "1/1",
  "ingress-mirror0": "switch1-cpu,unmodified",
  "ingress-mirror1": "switch1-cpu,unmodified",
  "mac-level-isolation": "true",
  "mirror-egress-if-ingress-mirrored": "false",
  "mirror-tx-on-mirror-port": "false",
  "mirrored-packet-drop-precedence": "green",
  "mirrored-packet-qos-priority": "0",
  "multicast-lookup-mode": "dst-ip-and-vid-for-ipv4",
  "name": "switch1",
  "override-existing-when-ufdb-full": "false",
  "type": "QCA-8511",
  "unicast-fdb-timeout": "5m",
  "unknown-vlan-lookup-mode": "svl",
  "use-cvid-in-one2one-vlan-lookup": "true",
  "use-svid-in-one2one-vlan-lookup": "false",
  "vlan-uses": "mirror0"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841835/CRS1xx+and+2xx+series+switches#CRS1xxand2xxseriesswitches-GlobalSettings
func ResourceInterfaceEthernetSwitchCrs() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch"),
		MetaId:           PropId(Id),

		"bridge_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The bridge type defines which VLAN tag is used as Lookup-VID. Lookup-VID serves as the VLAN " +
				"key for all VLAN-based lookups.",
			ValidateFunc:     validation.StringInSlice([]string{"customer-vid-used-as-lookup-vid", "service-vid-used-as-lookup-vid"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bypass_ingress_port_policing_for": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Protocols that are excluded from Ingress Port Policing. (arp, dhcpv4, dhcpv6, eapol, igmp, " +
				"mld, nd, pppoe-discovery, ripv1).",
		},
		"bypass_l2_security_check_filter_for": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Protocols that are excluded from Policy rule security check. (arp, dhcpv4, dhcpv6, eapol, " +
				"igmp, mld, nd, pppoe-discovery, ripv1).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bypass_vlan_ingress_filter_for": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Protocols that are excluded from Ingress VLAN filtering. These protocols are not dropped if " +
				"they have invalid VLAN. (arp, dhcpv4, dhcpv6,eapol, igmp, mld, nd, pppoe-discovery, ripv1).",
		},
		"drop_if_invalid_or_src_port_not_member_of_vlan_on_ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description:      "Ports that drop invalid and other port VLAN ID frames.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"drop_if_no_vlan_assignment_on_ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Ports which drop frames if no MAC-based, Protocol-based VLAN assignment or Ingress VLAN Translation " +
				"is applied.",
		},
		"egress_mirror_ratio": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Proportion of egress mirrored packets compared to all packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"egress_mirror0": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The first egress mirroring analyzer port or trunk and mirroring format:analyzer-configured " +
				"- The packet is the same as the packet to the destination. VLAN format is modified based on the VLAN " +
				"configurations of the analyzer port.modified - The packet is the same as the packet to the destination. " +
				"VLAN format is modified based on the VLAN configurations of the egress port.original - Traffic is mirrored " +
				"without any change to the original incoming packet format. But the service VLAN tag is stripped in the " +
				"edge port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"egress_mirror1": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The second egress mirroring analyzer port or trunk and mirroring format:analyzer-configured " +
				"- The packet is the same as the packet to the destination. VLAN format is modified based on the VLAN " +
				"configurations of the analyzer port.modified - The packet is the same as the packet to the destination. " +
				"VLAN format is modified based on the VLAN configurations of the egress port.original - Traffic is mirrored " +
				"without any change to the original incoming packet format. But the service VLAN tag is stripped in the " +
				"edge port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fdb_uses": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Analyzer port used for FDB-based mirroring.",
			ValidateFunc:     validation.StringInSlice([]string{"mirror0", "mirror1"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"forward_unknown_vlan": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to allow forwarding VLANs that are not members of the VLAN table.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ingress_mirror_ratio": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The proportion of ingress mirrored packets compared to all packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ingress_mirror0": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The first ingress mirroring analyzer port or trunk and mirroring format:analyzer-configured " +
				"- The packet is the same as the packet to the destination. VLAN format is modified based on the VLAN " +
				"configurations of the analyzer port.modified - The packet is the same as the packet to the destination. " +
				"VLAN format is modified based on the VLAN configurations of the egress port.original - Traffic is mirrored " +
				"without any change to the original incoming packet format. But the service VLAN tag is stripped in the " +
				"edge port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ingress_mirror1": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The second ingress mirroring analyzer port or trunk and mirroring format:analyzer-configured " +
				"- The packet is the same as the packet to the destination. VLAN format is modified based on the VLAN " +
				"configurations of the analyzer port.modified - The packet is the same as the packet to the destination. " +
				"VLAN format is modified based on the VLAN configurations of the egress port.original - Traffic is mirrored " +
				"without any change to the original incoming packet format. But the service VLAN tag is stripped in the " +
				"edge port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mac_level_isolation": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Globally enables or disables MAC level isolation. Once enabled, the switch will check the " +
				"source and destination MAC address entries and their isolation-profile from the unicast forwarding table. " +
				"By default, the switch will learn MAC addresses and place them into a promiscuous isolation profile. " +
				"Other isolation profiles can be used when creating static unicast entries. If the source or destination " +
				"MAC address is located on a promiscuous isolation profile, the packet is forwarded. If both source and " +
				"destination MAC addresses are located on the same community1 or community2 isolation profile, the packet " +
				"is forwarded. The packet is dropped when the source and destination MAC address isolation profile is " +
				"isolated, or when the source and destination MAC address isolation profiles are from different communities " +
				"(e.g. source MAC address is community1 and destination MAC address is community2). When MAC level isolation " +
				"is globally disabled, the isolation is bypassed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirror_egress_if_ingress_mirrored": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When a packet is applied to both ingress and egress mirroring, only ingress mirroring is performed " +
				"on the packet, if this setting is disabled. If this setting is enabled both mirroring types are applied.",
		},
		"mirror_tx_on_mirror_port": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "",
		},
		"mirrored_packet_drop_precedence": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Remarked drop precedence in mirrored packets. This QoS attribute is used for mirrored packet " +
				"enqueuing or dropping.",
			ValidateFunc:     validation.StringInSlice([]string{"drop", "green", "red", "yellow"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirrored_packet_qos_priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Remarked priority in mirrored packets.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multicast_lookup_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Lookup mode for IPv4 multicast bridging.dst-mac-and-vid-always - For all packet types lookup " +
				"key is the destination MAC and VLAN ID.dst-ip-and-vid-for-ipv4 - For IPv4 packets lookup key is the " +
				"destination IP and VLAN ID. For other packet types, the lookup key is the destination MAC and VLAN ID.",
			ValidateFunc:     validation.StringInSlice([]string{"dst-ip-and-vid-for-ipv4", "dst-mac-and-vid-always"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the switch."),
		"override_existing_when_ufdb_full": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enable or disable to override existing entry which has the lowest aging value when UFDB is " +
				"full.",
		},
		"unicast_fdb_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Timeout for Unicast FDB entries.",
			DiffSuppressFunc: TimeEqual,
		},
		"unknown_vlan_lookup_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Lookup and learning mode for packets with invalid VLAN.",
			ValidateFunc:     validation.StringInSlice([]string{"ivl", "svl"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_cvid_in_one2one_vlan_lookup": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to use customer VLAN ID for 1:1 VLAN switching lookup.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_svid_in_one2one_vlan_lookup": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to use service VLAN ID for 1:1 VLAN switching lookup.",
		},
		"vlan_uses": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Analyzer port used for VLAN-based mirroring.",
			ValidateFunc:     validation.StringInSlice([]string{"mirror0", "mirror1"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		Description: "Resource for managing CRS (Cloud Router Switch) series device properties.",

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
