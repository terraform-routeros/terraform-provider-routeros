package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
[
   {
        ".id": "*1",
        "add-default-route": "false",
        "dhcp-options": "",
        "dhcp-server-v6": "fe80::",
        "disabled": "false",
        "duid": "0x000003434343443",
        "interface": "if-name",
        "invalid": "false",
        "pool-name": "blacknight-pub-addr",
        "pool-prefix-length": "64",
        "prefix": "2a01:----:/56, 6d16h56m8s",
        "prefix-hint": "::/0",
        "request": "prefix",
        "status": "bound",
        "use-peer-dns": "true"
    }
]
*/

// ResourceIPv6DhcpClient https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPv6Client
func ResourceIPv6DhcpClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to add default IPv6 route after a client connects.",
		},
		"address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IPv6 address, which is assigned to DHCPv6 Client from the Server.",
		},
		KeyComment: PropCommentRw,
		"default_route_distance": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Distance of default route. Applicable if add-default-route is set to yes.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dhcp_options": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Options that are sent to the DHCP server.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"dhcp_server_v6": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IPv6 address of the DHCP server",
		},
		KeyDisabled: PropDisabledRw,
		"duid": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Auto-generated DUID that is sent to the server. DUID is generated using one of the MAC " +
				"addresses available on the router.",
		},
		KeyDynamic: PropDynamicRo,
		"expires_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A time when the IPv6 prefix expires (specified by the DHCPv6 server).",
		},
		"gateway": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the gateway which is assigned by DHCP server.",
		},
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"pool_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the IPv6 pool in which received IPv6 prefix will be added",
		},
		"pool_prefix_length": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Prefix length parameter that will be set for IPv6 pool in which received IPv6 prefix is " +
				"added. Prefix length must be greater than the length of the received prefix, otherwise, prefix-length " +
				"will be set to received prefix length + 8 bits.",
			ValidateFunc: validation.IntBetween(0, 128),
		},
		"prefix": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows received IPv6 prefix from DHCPv6-PD server",
		},
		"prefix_hint": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Description:      "Include a preferred prefix length.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rapid_commit": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable DHCP rapid commit (fast address assignment)",
		},
		"request": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "To choose if the DHCPv6 request will ask for the address, info or the IPv6 prefix.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"info", "address", "prefix"}, false),
			},
		},
		"script": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Run this script on the DHCP-client status change. Available variables:" +
				"\n  * pd-valid - if the prefix is acquired by the client;" +
				"\n  * pd-prefix - the prefix acquired by the client if any;" +
				"\n  * na-valid - if the address is acquired by the client;" +
				"\n  * na-address - the address acquired by the client if any." +
				"\n  * options - array of received options (only ROSv7)",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows the status of DHCPv6 Client:" +
				"\n  * stopped - dhcpv6 client is stopped" +
				"\n  * searching - sending `solicit` and trying to get `advertise`  Shows actual (resolved) gateway and " +
				"interface that will be used for packet forwarding.requesting - sent `request` waiting for `reply`" +
				"\n  * bound - received `reply`. Prefix assigned." +
				"\n  * renewing - sent `renew`, waiting for `reply`" +
				"\n  * rebinding - sent `rebind`, waiting for `reply`" +
				"\n  * error - reply was not received in time or some other error occurred." +
				"\n  * stopping - sent `release`",
		},
		"use_interface_duid": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Specifies the MAC address of the specified interface as the DHCPv6 client DUID.",
		},
		"use_peer_dns": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to accept the DNS settings advertised by the IPv6 DHCP Server.",
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
