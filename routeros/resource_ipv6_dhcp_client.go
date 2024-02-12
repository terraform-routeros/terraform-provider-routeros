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
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-client/"),
		MetaId:           PropId(Id),

		"add-default-route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to add default IPv6 route after a client connects.",
			Default:     true,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"duid": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Auto-generated DUID that is sent to the server." +
				"DUID is generated using one of the MAC addresses available on the router.",
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The interface on which the DHCPv6 client will be running.",
		},
		"pool_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the IPv6 pool in which received IPv6 prefix will be added",
		},
		"pool_prefix_length": {
			Type:     schema.TypeInt,
			Computed: false,
			Required: true,
			Description: "Prefix length parameter that will be set for IPv6 pool in which received IPv6 prefix is added." +
				" Prefix length must be greater than the length of the received prefix, otherwise, prefix-length will be set to received prefix length + 8 bits.",
			ValidateFunc: validation.IntBetween(0, 128),
		},
		"prefix": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows received IPv6 prefix from DHCPv6-PD server",
		},
		"prefix_hint ": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Include a preferred prefix length.",
			ValidateFunc: validation.IntBetween(0, 128),
		},
		"request": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "To choose if the DHCPv6 request will ask for the address or the IPv6 prefix, or both.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Shows the status of DHCPv6 Client:" +
				"stopped - dhcpv6 client is stopped" +
				"searching - sending \"solicit\" and trying to get \"advertise\"  Shows actual (resolved) gateway and interface that will be used for packet forwarding.requesting - sent \"request\" waiting for \"reply\"" +
				"bound - received \"reply\". Prefix assigned. " +
				"renewing - sent \"renew\", waiting for \"reply\" " +
				"rebinding - sent \"rebind\", waiting for \"reply\" " +
				"error - reply was not received in time or some other error occurred. " +
				"stopping - sent \"release\"",
		},
		"script": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Run this script on the DHCP-client status change. Available variables:" +
				"pd-valid - if the prefix is acquired by the client;" +
				"pd-prefix - the prefix acquired by the client if any;" +
				"na-valid - if the address is acquired by the client;" +
				"na-address - the address acquired by the client if any." +
				"options - array of received options (only ROSv7)",
		},
		"use_peer_dns": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Routing table this route belongs to.",
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
