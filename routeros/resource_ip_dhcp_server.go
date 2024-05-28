package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServer https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-Leases
func ResourceDhcpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server"),
		MetaId:           PropId(Id),

		"add_arp": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to add dynamic ARP entry. ",
		},
		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IP pool, from which to take IP addresses for the clients. If set to static-only, then only " +
				"the clients that have a static lease (added in lease submenu) will be allowed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allow_dual_stack_queue": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and " +
				"DUID for identification. Requires IPv6 DHCP Server to have this option enabled as well to work properly.",
		},
		"always_broadcast": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Always send replies as broadcasts even if destination IP is known.",
		},
		"authoritative": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Option changes the way how a server responds to DHCP requests.",
			ValidateFunc:     validation.StringInSlice([]string{"after-10sec-delay", "after-2sec-delay", "yes", "no"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bootp_lease_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Accepts two predefined options or time value: * forever - lease never expires " +
				"* lease-time - use time from lease-time parameter",
		},
		"bootp_support": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Support for BOOTP clients.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "static", "dynamic"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"client_mac_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Specifies whether to limit specific number of clients per single MAC address.",
		},
		"conflict_detection": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allows to disable/enable conflict detection. If option is enabled, then whenever server tries " +
				"to assign a lease it will send ICMP and ARP messages to detect whether such address in the network " +
				"already exist. If any of above get reply address is considered already used. Conflict detection must " +
				"be disabled when any kind of DHCP client limitation per port or per mac is used.",
		},
		KeyComment: PropCommentRw,
		"delay_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. " +
				"If set to none - there is no threshold (all DHCP packets are processed).",
		},
		"dhcp_option_set": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Use custom set of DHCP options defined in option sets menu.",
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"insert_queue_before": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify where to place dynamic simple queue entries for static DCHP leases with rate-limit parameter set.",
		},
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"lease_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A script that will be executed after a lease is assigned or de-assigned.",
		},
		"lease_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time that a client may use the assigned address. The client will try to renew this " +
				"address after half of this time and will request a new address after the time limit expires.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropNameForceNewRw,
		"parent_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
		"relay": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The IP address of the relay this DHCP server.",
			ValidateFunc: ValidationIpAddress,
		},
		"src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The address which the DHCP client must send requests to in order to renew an IP address lease.",
			ValidateFunc: ValidationIpAddress,
		},
		"use_framed_as_classless": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client.",
		},
		"use_radius": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to use RADIUS server.",
			ValidateFunc:     validation.StringInSlice([]string{"yes", "no", "accounting"}, false),
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

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceDhcpServerV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
