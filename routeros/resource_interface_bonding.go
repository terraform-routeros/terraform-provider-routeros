package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
".about": "802.3ad mode must use mii link-monitoring",
".id": "*B",
"arp": "enabled",
"arp-interval": "100ms",
"arp-ip-targets": "0.0.2.0",
"arp-timeout": "auto",
"disabled": "false",
"down-delay": "10ms",
"lacp-rate": "1sec",
"lacp-user-key": "123",
"link-monitoring": "arp",
"mac-address": "52:54:00:12:34:58",
"mii-interval": "100ms",
"min-links": "1",
"mlag-id": "321",
"mode": "802.3ad",
"mtu": "1500",
"name": "bonding1",
"primary": "none",
"running": "true",
"slaves": "ether3,ether4",
"transmit-hash-policy": "layer-2-and-3",
"up-delay": "10ms"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Bonding#Bonding-PropertyDescription
func ResourceInterfaceBonding() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bonding"),
		MetaId:           PropId(Id),

		"arp": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "enabled",
			Description: "Address Resolution Protocol for the interface. disabled - the interface will not use ARP " +
				"enabled - the interface will use ARP proxy-arp - the interface will use the ARP proxy " +
				"feature reply-only -the interface will only reply to requests originated from matching " +
				"IPaddress/MAC address combinations which are entered as static entries inthe '/ip " +
				"arp' table. No dynamic entries will be automatically stored inthe '/ip arp' table. " +
				"Therefore for communications to be successful, avalid static entry must already exist.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled", "proxy-arp", "reply-only"}, false),
		},
		"arp_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "100ms",
			Description:      "Time in milliseconds defines how often to monitor ARP requests.",
			DiffSuppressFunc: TimeEquall,
		},
		"arp_ip_targets": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IP target address which will be monitored if link-monitoring is set to arp. You can " +
				"specify multiple IP addresses, separated by a comma.",
		},
		KeyArpTimeout: PropArpTimeoutRw,
		KeyComment:    PropCommentRw,
		KeyDisabled:   PropDisabledRw,
		"down_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0",
			Description: "If a link failure has been detected, the bonding interface is disabled for a down-delay " +
				"time. The value should be a multiple of mii-interval, otherwise, it will be rounded down " +
				"to the nearest value. This property only has an effect when link-monitoring is set to mii.",
			DiffSuppressFunc: TimeEquall,
		},
		"forced_mac_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Bydefault, the bonding interface will use the MAC address of the firstselected slave " +
				"interface. This property allows to configure static MACaddress for the bond interface " +
				"(all zeros, broadcast or multicastaddresses will not apply). RouterOS will " +
				"automatically change the MACaddress for slave interfaces and it will be visible in " +
				"/interface ethernet configuration export.",
		},
		"lacp_rate": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "30secs",
			Description: "LinkAggregation Control Protocol rate specifies how often to exchange withLACPDUs " +
				"between bonding peers. Used to determine whether a link is up orother changes have " +
				"occurred in the network. LACP tries to adapt tothese changes providing failover.",
			ValidateFunc: validation.StringInSlice([]string{"1sec", "30secs"}, false),
		},
		"lacp_user_key": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Specifiesthe upper 10 bits of the port key. The lower 6 bits are automatically" +
				"assigned based on individual port link speed and duplex. The setting isavailable only " +
				"since RouterOS v7.3.",
			ValidateFunc: validation.IntBetween(0, 1023),
		},
		"link_monitoring": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "mii",
			Description: "Method to use for monitoring the link (whether it is up or down) arp - uses Address " +
				"Resolution Protocol to determine whether the remote interface is reachable mii - uses " +
				"Media Independent Interface to determine link status. Link status determination relies " +
				"on the device driver. none - no method for link monitoring is used. Note: some bonding " +
				"modes require specific link monitoring to work properly.",
			ValidateFunc: validation.StringInSlice([]string{"arp", "mii", "none"}, false),
		},
		KeyMacAddress: PropMacAddressRo,
		"min_links": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "How many active slave links needed for bonding to become active.",
			ValidateFunc: validation.IntAtLeast(0),
		},
		"mii_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "100ms",
			Description: "How often to monitor the link for failures (the parameter used only if link-monitoring " +
				"is mii)",
			DiffSuppressFunc: TimeEquall,
		},
		"mlag_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "ChangesMLAG ID for bonding interface. The same MLAG ID should be used on bothpeer " +
				"devices to successfully create a single MLAG. See more details on MLAG .",
			ValidateFunc: validation.IntAtLeast(0),
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "balance-rr",
			Description: "Specifies one of the bonding policies:\n  * 802.3ad -IEEE 802.3ad dynamic link aggregation. " +
				"In this mode, the interfaces areaggregated in a group where each slave shares the same " +
				"speed. Itprovides fault tolerance and load balancing. Slave selection foroutgoing " +
				"traffic is done according to the transmit-hash-policy\n  * active-backup - provides " +
				"link backup. Only one slave can be active at a time. Another slave only becomes active, " +
				"if the first one fails.\n  * balance-alb - adaptive load balancing. The same as " +
				"balance-tlb but received traffic is also balanced. The device driver should have support " +
				"for changing it's MAC address.\n  * balance-rr -round-robin load balancing. Slaves in " +
				"a bonding interface will transmitand receive data in sequential order. It provides " +
				"load balancing andfault tolerance.\n  * balance-tlb -Outgoing traffic is " +
				"distributed according to the current load on eachslave. Incoming traffic is not " +
				"balanced and is received by the currentslave. If receiving slave fails, then another " +
				"slave takes the MACaddress of the failed slave.\n  * balance-xor - Transmit based on " +
				"the selected transmit-hash-policy. This mode provides load balancing and fault " +
				"tolerance.\n  * broadcast -Broadcasts the same data on all interfaces at once. This " +
				"provides faulttolerance but slows down traffic throughput on some slow machines.",
			ValidateFunc: validation.StringInSlice(
				[]string{
					"802.3ad",
					"active-backup",
					"balance-alb",
					"balance-rr",
					"balance-tlb",
					"balance-xor",
					"broadcast",
				},
				false,
			),
		},
		KeyMtu: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
			Description: "MaximumTransmit Unit in bytes. Must be smaller or equal to the smallest L2MTUvalue " +
				"of a bonding slave. L2MTU of a bonding interface is determined bythe lowest L2MTU " +
				"value among its slave interfaces.",
			ValidateFunc: validation.IntBetween(64, 65535),
		},
		KeyName: PropName("Name of the bonding interface."),
		"primary": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
			Description: "Controlsthe primary interface between active slave ports, works only for" +
				"active-backup, balance-tlb and balance-alb modes. For active-backupmode, it controls " +
				"which running interface is supposed to send andreceive the traffic. For balance-tlb " +
				"mode, it controls which runninginterface is supposed to receive all the traffic, but " +
				"for balance-albmode, it controls which interface is supposed to receive the unbalanced " +
				" traffic (the non-IPv4 traffic). When none of the interfaces are selectedas primary, " +
				"device will automatically select the interface that isconfigured as the first one.",
		},
		KeyRunning: PropRunningRo,
		"slaves": {
			Type:     schema.TypeSet,
			Required: true,
			Description: "At least two ethernet-like interfaces separated by a comma, which will be used for " +
				"bonding",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"up_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0",
			Description: "If a link has been brought up, the bonding interface is disabled for up-delay time and " +
				"after this time it is enabled. The value should be a multiple of mii-interval , " +
				"otherwise, it will be rounded down to the nearest value. This property only has an " +
				"effect when link-monitoring is set to mii.",
			DiffSuppressFunc: TimeEquall,
		},
		"transmit_hash_policy": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "layer-2",
			Description: "Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad " +
				"modes:\n  * layer-2 -Uses XOR of hardware MAC addresses to generate the hash. This algorithm " +
				" will place all traffic to a particular network peer on the same slave.This algorithm " +
				"is 802.3ad compliant.\n  * layer-2-and-3 -This policy uses a combination of layer2 and " +
				"layer3 protocolinformation to generate the hash. Uses XOR of hardware MAC addresses " +
				"andIP addresses to generate the hash. This algorithm will place alltraffic to a " +
				"particular network peer on the same slave. For non-IPtraffic, the formula is the same " +
				"as for the layer2 transmit hash policy.This policy is intended to provide a more " +
				"balanced distribution oftraffic than layer2 alone, especially in environments where a " +
				"layer3gateway device is required to reach most destinations. This algorithm is" +
				"802.3ad compliant.\n  * layer-3-and-4 - This policyuses upper layer protocol information, " +
				"when available, to generate thehash. This allows for traffic to a particular network " +
				"peer to spanmultiple slaves, although a single connection will not span multiple" +
				"slaves. For fragmented TCP or UDP packets and all other IP protocoltraffic, the source " +
				"and destination port information is omitted. Fornon-IP traffic, the formula is the " +
				"same as for the layer2 transmit hashpolicy. This algorithm is not fully 802.3ad " +
				"compliant.",
			ValidateFunc: validation.StringInSlice([]string{"layer-2", "layer-2-and-3", "layer-3-and-4"}, false),
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
