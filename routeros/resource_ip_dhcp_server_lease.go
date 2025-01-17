package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServerLease https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Server
func ResourceDhcpServerLease() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/lease"),
		MetaId:           PropId(Id),

		"active_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the machine currently holding the DHCP lease.",
		},
		"active_client_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Actual client-id of the client.",
		},
		"active_hostname": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The hostname of the machine currently holding the DHCP lease.",
		},
		"active_mac_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The MAC address of of the machine currently holding the DHCP lease.",
		},
		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IP address of the DHCP lease to be created.",
		},
		"address_lists": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Address list to which address will be added if lease is bound.",
		},
		"active_server": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Actual dhcp server, which serves this client.",
		},
		"agent_circuit_id": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Circuit ID of DHCP relay agent. If each character should be valid ASCII text symbol or else " +
				"this value is displayed as hex dump.",
		},
		"agent_remote_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Remote ID, set by DHCP relay agent.",
		},
		"allow_dual_stack_queue": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address " +
				"and DUID for identification.",
		},
		"always_broadcast": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Send all replies as broadcasts.",
		},
		"class_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Class ID of the client. This option is available in RouterOS starting from version 7.16.",
		},
		"block_access": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to block access for this DHCP client (true|false).",
		},
		"blocked": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the lease is blocked.",
		},
		// bridge_port ?
		"client_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If specified, must match DHCP 'client identifier' option of the request.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dhcp_option": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Add additional DHCP options.",
		},
		"dhcp_option_set": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Add additional set of DHCP options.",
		},
		"dynamic": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the dhcp lease is static or dynamic. Dynamic leases are not guaranteed to continue to be assigned to that specific device.",
		},
		"expires_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time until lease expires.",
		},
		"host_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The hostname of the device",
		},
		"insert_queue_before": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify where to place dynamic simple queue entries for static DCHP leases with " +
				"rate-limit parameter set.",
		},
		"last_seen": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"lease_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Time that the client may use the address. If set to 0s lease will never expire.",
		},
		"mac_address": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "The MAC addreess of the DHCP lease to be created.",
			ValidateFunc:     validation.IsMACAddress,
			DiffSuppressFunc: MacAddressEqual,
		},
		"radius": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows if this dynamic lease is authenticated by RADIUS or not.",
		},
		"rate_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. " +
				"Requires the lease to be static.",
		},
		"server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Server name which serves this client.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Lease status.",
		},
		"src_mac_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Source MAC address.",
		},
		"use_src_mac": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When this option is set server uses source MAC address instead of received CHADDR to " +
				"assign address.",
		},
	}
	return &schema.Resource{
		Description: "Creates a DHCP lease on the mikrotik device.",

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
