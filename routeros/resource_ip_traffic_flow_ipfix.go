package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "bytes": "true",
  "dst-address": "true",
  "dst-address-mask": "true",
  "dst-mac-address": "true",
  "dst-port": "true",
  "first-forwarded": "true",
  "gateway": "true",
  "icmp-code": "true",
  "icmp-type": "true",
  "igmp-type": "true",
  "in-interface": "true",
  "ip-header-length": "true",
  "ip-total-length": "true",
  "ipv6-flow-label": "true",
  "is-multicast": "true",
  "last-forwarded": "true",
  "nat-dst-address": "true",
  "nat-dst-port": "true",
  "nat-events": "false",
  "nat-src-address": "true",
  "nat-src-port": "true",
  "out-interface": "true",
  "packets": "true",
  "protocol": "true",
  "src-address": "true",
  "src-address-mask": "true",
  "src-mac-address": "true",
  "src-port": "true",
  "sys-init-time": "true",
  "tcp-ack-num": "true",
  "tcp-flags": "true",
  "tcp-seq-num": "true",
  "tcp-window-size": "true",
  "tos": "true",
  "ttl": "true",
  "udp-length": "true"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/21102653/Traffic+Flow#TrafficFlow-IPFIX
func ResourceIpTrafficFlowIpfix() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/ip/traffic-flow/ipfix"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("nat_events"),

		"bytes": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Total number of bytes processed in the flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The destination IP address of the flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_address_mask": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Network mask for the destination address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_mac_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Destination MAC address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Destination port number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"first_forwarded": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Timestamp of the first packet forwarded in a flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"gateway": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address of the gateway through which the flow was routed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"icmp_code": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "ICMP code for error messaging and operational information.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"icmp_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of ICMP message, important for diagnostic messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"igmp_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of Internet Group Management Protocol operation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"in_interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interface through which packets of the flow are received.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ip_header_length": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Length of the IP header.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ip_total_lenght": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Length of the IP packet in bytes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipv6_flow_label": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Label field from an IPv6 header, used to classify flows.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"is_multicast": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Indicates whether the flow is a multicast flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"last_forwarded": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Timestamp of the last packet forwarded in a flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nat_dst_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Translated destination IP address by NAT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nat_dst_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Translated destination port number by NAT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nat_events": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Events related to Network Address Translation for the flow.",
		},
		"nat_src_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Translated source IP address by NAT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nat_src_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Translated source port number by NAT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"out_interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interface through which packets of the flow are sent out.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"packets": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Number of packets processed in the flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"protocol": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Protocol number (e.g., TCP, UDP, ICMP).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The source IP address of the flow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address_mask": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Network mask for the source address, useful in summarizing data.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_mac_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Source MAC address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Source port number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sys_init_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "System initialization time, can be used for timing analysis.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_ack_num": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Acknowledgment number in a TCP connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_flags": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Flags from the TCP header (e.g., SYN, ACK).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_seq_num": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sequence number in a TCP connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_window_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Window size in a TCP connection, indicating the scale of received data buffering.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tos": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Type of Service field in the IP header, indicating priority and handling of the packet.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ttl": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time To Live for the packet, decremented by each router to prevent infinite loops.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"udp_length": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Length of the UDP payload.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
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
