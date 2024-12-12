package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "file-limit": "1000",
  "file-name": "",
  "filter-cpu": "",
  "filter-direction": "any",
  "filter-dst-ip-address": "",
  "filter-dst-ipv6-address": "",
  "filter-dst-mac-address": "",
  "filter-dst-port": "",
  "filter-interface": "",
  "filter-ip-address": "",
  "filter-ip-protocol": "",
  "filter-ipv6-address": "",
  "filter-mac-address": "",
  "filter-mac-protocol": "",
  "filter-operator-between-entries": "or",
  "filter-port": "",
  "filter-size": "",
  "filter-src-ip-address": "",
  "filter-src-ipv6-address": "",
  "filter-src-mac-address": "",
  "filter-src-port": "",
  "filter-stream": "false",
  "filter-vlan": "",
  "memory-limit": "100",
  "memory-scroll": "true",
  "only-headers": "false",
  "quick-rows": "20",
  "quick-show-frame": "false",
  "running": "false",
  "streaming-enabled": "false",
  "streaming-server": "0.0.0.0:37008"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Packet+Sniffer
func ResourceToolSniffer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/sniffer"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("quick_rows", "quick_show_frame", "show_frame", "enabled"),

		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Start packet capture.",
		},
		"file_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "File size limit. Sniffer will stop when a limit is reached.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"file_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the file where sniffed packets will be saved.",
		},
		"filter_cpu": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "CPU core used as a filter.",
		},
		"filter_direction": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which direction filtering will be applied.",
			ValidateFunc:     validation.StringInSlice([]string{"any", "rx", "tx"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"filter_dst_mac_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 MAC destination addresses and MAC address masks used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsMACAddress,
			},
			MaxItems: 16,
		},
		"filter_dst_ip_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IP destination addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			MaxItems: 16,
		},
		"filter_dst_ipv6_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IPv6 destination addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv6Address,
			},
			MaxItems: 16,
		},
		"filter_dst_port": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Up to 16 comma-separated destination ports used as a filter. A list of predefined port names " +
				"is also available, like ssh and telnet.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MaxItems: 16,
		},
		"filter_interface": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Interface name on which sniffer will be running. all indicates that the sniffer will sniff " +
				"packets on all interfaces.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MaxItems: 16,
		},
		"filter_ip_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IP addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			MaxItems: 16,
		},
		"filter_ipv6_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IPv6 addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv6Address,
			},
			MaxItems: 16,
		},
		"filter_ip_protocol": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Up to 16 comma-separated IP/IPv6 protocols used as a filter. IP protocols (instead of protocol " +
				"names, protocol numbers can be used):" +
				"\n  * ipsec-ah - IPsec AH protocol" +
				"\n  * ipsec-esp - IPsec ESP protocol" +
				"\n  * ddp - datagram delivery protocol" +
				"\n  * egp - exterior gateway protocol" +
				"\n  * ggp - gateway-gateway protocol" +
				"\n  * gre - general routing encapsulation" +
				"\n  * hmp - host monitoring protocol" +
				"\n  * idpr-cmtp - idpr control message transport" +
				"\n  * icmp - internet control message protocol" +
				"\n  * icmpv6 - internet control message protocol v6" +
				"\n  * igmp - internet group management protocol" +
				"\n  * ipencap - ip encapsulated in ip" +
				"\n  * ipip - ip encapsulation" +
				"\n  * encap - ip encapsulation" +
				"\n  * iso-tp4 - iso transport protocol class 4" +
				"\n  * ospf - open shortest path first" +
				"\n  * pup - parc universal packet protocol" +
				"\n  * pim - protocol independent multicast" +
				"\n  * rspf - radio shortest path first" +
				"\n  * rdp - reliable datagram protocol" +
				"\n  * st - st datagram mode" +
				"\n  * tcp - transmission control protocol" +
				"\n  * udp - user datagram protocol" +
				"\n  * vmtp versatile message transport" +
				"\n  * vrrp - virtual router redundancy protocol" +
				"\n  * xns-idp - xerox xns idp" +
				"\n  * xtp - xpress transfer protocol",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationValInSlice([]string{}, false, true),
			},
			MaxItems: 16,
		},
		"filter_mac_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 MAC addresses and MAC address masks used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsMACAddress,
			},
			MaxItems: 16,
		},
		"filter_mac_protocol": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Up to 16 comma separated entries used as a filter. Mac protocols (instead of protocol names, " +
				"protocol number can be used):" +
				"\n  * 802.2 - 802.2 Frames (0x0004)" +
				"\n  * arp - Address Resolution Protocol (0x0806)" +
				"\n  * homeplug-av - HomePlug AV MME (0x88E1)" +
				"\n  * ip - Internet Protocol version 4 (0x0800)" +
				"\n  * ipv6 - Internet Protocol Version 6 (0x86DD)" +
				"\n  * ipx - Internetwork Packet Exchange (0x8137)" +
				"\n  * lldp - Link Layer Discovery Protocol (0x88CC)" +
				"\n  * loop-protect - Loop Protect Protocol (0x9003)" +
				"\n  * mpls-multicast - MPLS multicast (0x8848)" +
				"\n  * mpls-unicast - MPLS unicast (0x8847)" +
				"\n  * packing-compr - Encapsulated packets with compressed IP packing (0x9001)" +
				"\n  * packing-simple - Encapsulated packets with simple IP packing (0x9000)" +
				"\n  * pppoe - PPPoE Session Stage (0x8864)" +
				"\n  * pppoe-discovery - PPPoE Discovery Stage (0x8863)" +
				"\n  * rarp - Reverse Address Resolution Protocol (0x8035)" +
				"\n  * service-vlan - Provider Bridging (IEEE 802.1ad) & Shortest Path Bridging IEEE 802.1aq (0x88A8)" +
				"\n  * vlan - VLAN-tagged frame (IEEE 802.1Q) and Shortest Path Bridging IEEE 802.1aq with NNI compatibility (0x8100)",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MaxItems: 16,
		},
		"filter_operator_between_entries": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Changes the logic for filters with multiple entries.",
			ValidateFunc:     validation.StringInSlice([]string{"and", "or"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"filter_port": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Up to 16 comma-separated ports used as a filter. A list of predefined port names is also available, " +
				"like ssh and telnet.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MaxItems: 16,
		},
		// Type attribute string instead of number because MT returns an empty string value.
		"filter_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Filters packets of specified size or size range in bytes.",
		},
		"filter_src_mac_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 MAC source addresses and MAC address masks used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsMACAddress,
			},
			MaxItems: 16,
		},
		"filter_src_ip_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IP source addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			MaxItems: 16,
		},
		"filter_src_ipv6_address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 IPv6 source addresses used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv6Address,
			},
			MaxItems: 16,
		},
		"filter_src_port": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Up to 16 comma-separated source ports used as a filter. A list of predefined port names is " +
				"also available, like ssh and telnet.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			MaxItems: 16,
		},
		"filter_stream": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Sniffed packets that are devised for the sniffer server are ignored.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"filter_vlan": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Up to 16 VLAN IDs used as a filter.",
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4095),
			},
		},
		"memory_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Memory amount used to store sniffed data.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"memory_scroll": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to rewrite older sniffed data when the memory limit is reached.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"only_headers": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Save in the memory only the packet's headers, not the whole packet.",
		},
		KeyRunning: PropRunningRo,
		"streaming_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Defines whether to send sniffed packets to the streaming server.",
		},
		"streaming_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Tazmen Sniffer Protocol (TZSP) stream receiver.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			diags := SystemResourceCreateUpdate(ctx, resSchema, d, m)
			if diags.HasError() {
				return diags
			}

			setSnifferState(ctx, resSchema, d, m)

			return SystemResourceRead(ctx, resSchema, d, m)
		},

		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			if d := SystemResourceRead(ctx, resSchema, d, m); d.HasError() {
				return d
			}
			d.Set("enabled", d.Get(KeyRunning).(bool))

			return nil
		},

		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			stopSniffer(ctx, resSchema, d, m)

			diags := SystemResourceCreateUpdate(ctx, resSchema, d, m)
			if diags.HasError() {
				return diags
			}
			setSnifferState(ctx, resSchema, d, m)

			return SystemResourceRead(ctx, resSchema, d, m)
		},

		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			stopSniffer(ctx, resSchema, d, m)

			return SystemResourceDelete(ctx, resSchema, d, m)
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}

func setSnifferState(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.Get("enabled").(bool) {
		return startSniffer(ctx, s, d, m)
	}
	return stopSniffer(ctx, s, d, m)
}

func startSniffer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Start sniffer.
	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/start"
	}

	err := m.(Client).SendRequest(crudStart, resUrl, MikrotikItem{}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func stopSniffer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Stop sniffer.
	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/stop"
	}

	err := m.(Client).SendRequest(crudStop, resUrl, MikrotikItem{}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
