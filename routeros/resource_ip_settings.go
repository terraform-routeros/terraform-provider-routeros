package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "accept-redirects": "false",
  "accept-source-route": "false",
  "allow-fast-path": "true",
  "arp-timeout": "30s",
  "icmp-errors-use-inbound-interface-address": "false",
  "icmp-rate-limit": "10",
  "icmp-rate-mask": "0x1818",
  "ip-forward": "true",
  "ipv4-fast-path-active": "true",
  "ipv4-fast-path-bytes": "0",
  "ipv4-fast-path-packets": "0",
  "ipv4-fasttrack-active": "false",
  "ipv4-fasttrack-bytes": "0",
  "ipv4-fasttrack-packets": "0",
  "ipv4-multipath-hash-policy": "l3",
  "max-neighbor-entries": "8192",
  "rp-filter": "no",
  "secure-redirects": "true",
  "send-redirects": "true",
  "tcp-syncookies": "false",
  "tcp-timestamps": "random-offset"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/103841817/IP+Settings
func ResourceIpSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/settings"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("ipv4_fast_path_active", "ipv4_fast_path_bytes", "ipv4_fast_path_packets",
			"ipv4_fasttrack_active", "ipv4_fasttrack_bytes", "ipv4_fasttrack_packets"),

		"accept_redirects": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to accept ICMP redirect messages. Typically should be enabled on the host and disabled " +
				"on routers.",
		},
		"accept_source_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to accept packets with the SRR option. Typically should be enabled on the router.",
		},
		"allow_fast_path": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Allows Fast Path.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"arp_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Sets Linux base_reachable_time (base_reachable_time_ms) on all interfaces that use ARP. " +
				"The initial validity of the ARP entry is picked from the interval [timeout/2 - 3*timeout/2] (default from " +
				"15s to 45s) after the neighbor was found. Can use postfix ms, s, m, h, d for milliseconds, seconds, " +
				"minutes, hours, or days. if no postfix is set then seconds (s) are used. The parameter means how long " +
				"a valid ARP record will be considered complete if no one communicates with the specific MAC/IP during " +
				"this time. The parameter does not represent a time when an ARP entry is removed from the ARP cache " +
				"(see max-neighbor-entries setting).",
			DiffSuppressFunc: TimeEqual,
		},
		"icmp_errors_use_inbound_interface_address": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If enabled, the ICMP error message reply will be sent with the source address equal to primary " +
				"address of the receiving interface that caused the error . This feature can be useful for complex network " +
				"debugging.",
		},
		"icmp_rate_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Limit the maximum rates for sending ICMP packets whose type matches icmp-rate-mask to specific " +
				"targets. `0` disables any limiting, other values indicate the minimum space between responses in milliseconds.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"icmp_rate_mask": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Mask made of ICMP types for which rates are being limited.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ip_forward": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enable/disable packet forwarding between interfaces. Resets all configuration parameters " +
				"to defaults according to RFC1812 for routers.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipv4_multipath_hash_policy": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IPv4 Hash policy used for ECMP routing in `/ip/settings` menu" +
				"\n  * l3 -- layer-3 hashing of src IP, dst IP" +
				"\n  * l3-inner -- layer-3 hashing or inner layer-3 hashing if available" +
				"\n  * l4 -- layer-4 hashing of src IP, dst IP, IP protocol, src port, dst port",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_neighbor_entries": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Sets Linux gc_thresh3. A maximum number of allowed neighbors in the ARP table. Since " +
				"`RouterOS version 7.1`, the default value depends on the installed amount of RAM. It is possible to set " +
				"a higher value than the default, but it increases the risk of out-of-memory condition." +
				"\nThe default values for certain RAM sizes:" +
				"\n  *  2048 for 64 MB," +
				"\n  *  4096 for 128 MB," +
				"\n  *  8192 for 256 MB," +
				"\n  *  16384 for 512 MB or higher." +
				"\nThe ARP cache stores ARP entries, and if some of these entries are incomplete, they can stay in the " +
				"cache for an indefinite period of time. This will only happen if the number of entries in the cache is " +
				"less than one-fourth of the maximum number allowed. The reason for this is to prevent the unnecessary " +
				"running of the garbage-collector when the ARP table is not close to being full.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"route_cache": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Disable or enable the Linux route cache. Note that disabling the route cache, will also " +
				"disable the fast path.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rp_filter": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Disables or enables source validation." +
				"\n  *  no - No source validation." +
				"\n  *  strict - Strict mode as defined in RFC3704 Strict Reverse Path. Each incoming packet is tested " +
				"against the FIB and if the interface is not the best reverse path the packet check will fail. By " +
				"default failed packets are discarded." +
				"\n  *  loose - Loose mode as defined in RFC3704 Loose Reverse Path. Each incoming packet's source " +
				"address is also tested against the FIB and if the source address is not reachable via any interface " +
				"the packet check will fail." +
				"\nThe current recommended practice in RFC3704 is to enable strict mode to prevent IP spoofing from DDoS " +
				"attacks. If using asymmetric routing or other complicated routing or VRRP, then the loose mode is recommended." +
				"\n`Warning`: strict mode does not work with routing tables",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secure_redirects": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Accept ICMP redirect messages only for gateways, listed in the default gateway list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_redirects": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to send ICMP redirects. Recommended to be enabled on routers.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_timestamps": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Parameter allows to enable/disable TCP timestamps or add random offset to TCP timestamp " +
				"(default behavior). Disabling timestamps completely may help to reduce spikes of performance drops.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_syncookies": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "end out syncookies when the syn backlog queue of a socket overflows. This is to prevent " +
				"the common 'SYN flood attack'. syncookies seriously violate TCP protocol, and disallow the use of TCP " +
				"extensions, which can result in serious degradation of some services (f.e. SMTP relaying), visible not " +
				"by you, but to your clients and relays, contacting you.",
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
