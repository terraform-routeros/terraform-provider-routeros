package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address-list": "",
    "bridge-learning": "default",
    "change-tcp-mss": "default",
    "default": "false",
	"insert-queue-before": "first",
    "local-address": "192.168.77.1",
    "name": "ovpn",
    "on-down": "",
    "on-up": "",
    "only-one": "default",
	"parent-queue": "none",
	"queue-type": "multi-queue-ethernet-default",
    "remote-address": "*2",
    "use-compression": "default",
    "use-encryption": "default",
    "use-ipv6": "yes",
    "use-mpls": "default",
    "use-upnp": "default"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/PPP+AAA#PPPAAA-UserProfiles
func ResourcePPPProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ppp/profile"),
		MetaId:           PropId(Id),

		"address_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Address list name to which ppp assigned (on server) or received (on client) address will " +
				"be added.",
		},
		"bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the bridge interface to which ppp interface will be added as a slave port. Both  " +
				"tunnel endpoints (server and client) must be in bridge in order to make  this work, see " +
				"more details on the BCP bridging manual.",
		},
		"bridge_horizon": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Used  split-horizon value for the dynamically created bridge port. Can be  used to " +
				"prevent bridging loops and isolate traffic. Set the same value  for a group of ports, to " +
				"prevent them from sending data to ports with  the same horizon value.",
			ValidateFunc: validation.IntAtLeast(1),
		},
		"bridge_learning": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Changes MAC learning behavior on the dynamically created bridge port: yes - enables MAC " +
				"learning no - disables MAC learning default - derive this value from the interface " +
				"default profile; same as yes if this is the interface default profile.",
			ValidateFunc: validation.StringInSlice([]string{"default", "no", "yes"}, false),
		},
		"bridge_path_cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Used  path cost for the dynamically created bridge port, used by STP/RSTP to  determine " +
				"the best path, used by MSTP to determine the best path between  regions. This property " +
				"has no effect when a bridge protocol-mode is set to none.",
			ValidateFunc: validation.IntAtLeast(0),
		},
		"bridge_port_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Used  priority for the dynamically created bridge port, used by STP/RSTP to  determine " +
				"the root port, used by MSTP to determine root port between  regions. This property has " +
				"no effect when a bridge protocol-mode is set  to none.",
			ValidateFunc: validation.IntBetween(0, 240),
		},
		"change_tcp_mss": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Modifies connection MSS settings (applies only for IPv4): yes - adjust connection MSS " +
				"value no - do not adjust connection MSS value default - derive this value from the " +
				"interface default profile; same as no if this is the interface default profile.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default"}, false),
		},
		KeyComment: PropCommentRw,
		KeyDefault: PropDefaultRo,
		"dhcpv6_pd_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the IPv6 pool which will be used by dynamically created DHCPv6-PD server when " +
				"client connects. [Read more >>](https://wiki.mikrotik.com/wiki/Manual:IPv6_PD_over_PPP)",
		},
		"dns_server": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "IP address of the DNS server that is supplied to ppp clients.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			MaxItems: 2,
		},
		"idle_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies  the amount of time after which the link will be terminated if there are  no " +
				"activity present. Timeout is not set by default.",
			DiffSuppressFunc: TimeEquall,
		},
		"incoming_filter": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Firewall  chain name for incoming packets. Specified chain gets control for each  packet " +
				"coming from the client. The ppp chain should be manually added  and rules with " +
				"action=jump jump-target=ppp should be added to other  relevant chains in order for this " +
				"feature to work. For more information  look at the examples section.",
		},
		"insert_queue_before": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit " +
				"parameter set.",
		},
		"interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface list name.",
		},
		"local_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Tunnel address or name of the pool from which address is assigned to ppp interface locally.",
		},
		KeyName: PropName("PPP profile name."),
		"on_up": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Execute script on user login-event. These are available variables that are accessible " +
				"for the event script:\n  * user\n  * local-address\n  * remote-address\n  * caller-id\n  * called-id\n  * interface.",
		},
		"on_down": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Execute script on user logging off. See on-up for more details.",
		},
		"only_one": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Defines whether a user is allowed to have more than one ppp session at a time yes - a " +
				"user is not allowed to have more than one ppp session at a time no - the user is allowed " +
				"to have more than one ppp session at a time default - derive this value from the " +
				"interface default profile; same as no if this is the interface default profile.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default"}, false),
		},
		"outgoing_filter": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Firewall  chain name for outgoing packets. The specified chain gets control for  each " +
				"packet going to the client. The PPP chain should be manually added  and rules with " +
				"action=jump jump-target=ppp should be added to other  relevant chains in order for this " +
				"feature to work. For more information  look at the Examples section.",
		},
		"parent_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of parent simple queue.",
		},
		"queue_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Queue types.",
		},
		"rate_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rate limitation in form of rx-rate[/tx-rate]  [rx-burst-rate[/tx-burst-rate] " +
				"[rx-burst-threshold[/tx-burst-threshold]  [rx-burst-time[/tx-burst-time] [priority] " +
				"[rx-rate-min[/tx-rate-min]]]] from the point of view of the router (so 'rx' is client " +
				"upload, and  'tx' is client download). All rates are measured in bits per second,  " +
				"unless followed by optional 'k' suffix (kilobits per second) or 'M'  suffix (megabits " +
				"per second). If tx-rate is not specified, rx-rate  serves as tx-rate too. The same " +
				"applies for tx-burst-rate,  tx-burst-threshold and tx-burst-time. If both " +
				"rx-burst-threshold and  tx-burst-threshold are not specified (but burst-rate is " +
				"specified),  rx-rate and tx-rate are used as burst thresholds. If both rx-burst-time  " +
				"and tx-burst-time are not specified, 1s is used as default. Priority  takes values 1..8, " +
				"where 1 implies the highest priority, but 8 - the  lowest. If rx-rate-min and " +
				"tx-rate-min are not specified rx-rate and  tx-rate values are used. The rx-rate-min and " +
				"tx-rate-min values can not  exceed rx-rate and tx-rate values.",
		},
		"remote_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Tunnel address or name of the pool from which address is assigned to remote ppp interface.",
		},
		"remote_ipv6_prefix_pool": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Assign prefix from IPv6 pool to the client and install corresponding IPv6 route.",
		},
		"session_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum time the connection can stay up. By default no time limit is set.",
			DiffSuppressFunc: TimeEquall,
		},
		"use_compression": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Specifies whether to use data compression or not. yes - enable data compression no - " +
				"disable data compression default - derive this value from the interface default profile; " +
				"same as no if this is the interface default profile This setting does not affect OVPN " +
				"tunnels.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default"}, false),
		},
		"use_encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Specifies whether to use data encryption or not. yes - enable data encryption no - " +
				"disable data encryption default - derive this value from the interface default profile; " +
				"same as no if this is the interface default profile require - explicitly requires " +
				"encryption This setting does not work on OVPN and SSTP tunnels.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default", "require"}, false),
		},
		"use_ipv6": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Specifies whether to allow IPv6. By default is enabled if IPv6 package is installed. yes " +
				"- enable IPv6 support no - disable IPv6 support default - derive this value from the " +
				"interface default profile; same as no if this is the interface default profile require - " +
				"explicitly requires IPv6 support.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default", "require"}, false),
		},
		"use_mpls": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
			Description: "Specifies whether to allow MPLS over PPP. yes - enable MPLS support no - disable MPLS " +
				"support default - derive this value from the interface default profile; same as no if " +
				"this is the interface default profile require - explicitly requires MPLS support",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default", "require"}, false),
		},
		"use_upnp": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "default",
			Description:  "Specifies whether to allow UPnP.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no", "default"}, false),
		},
		"wins_server": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "IP address of the WINS server to supply to Windows clients.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			MaxItems: 2,
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
