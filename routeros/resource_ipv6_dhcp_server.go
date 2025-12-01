package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address-pool": "ULA",
    "comment": "https://ula.ungleich.ch/random/",
    "dhcp-option": "dns",
    "disabled": "false",
    "duid": "0x00030001d401c330e280",
    "dynamic": "false",
    "interface": "span-bridge",
    "invalid": "false",
    "lease-time": "10m",
    "name": "server1",
    "preference": "255",
    "rapid-commit": "true",
    "route-distance": "1",
    "use-radius": "false"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPv6Server
func ResourceIpv6DhcpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-server"),
		MetaId:           PropId(Id),

		"address_pool": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IPv6 pool, from which to take IPv6 address for the clients. The prefix length of the pool " +
				"must be 128.",
			AtLeastOneOf:     []string{"address_pool", "prefix_pool"},
		},
		"address_lists": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Firewall address lists to which the allocated addresses and prefixes will be added if the lease is bound.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allow_dual_stack_queue": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Creates a single simple queue entry for both IPv4 and IPv6 addresses, and uses the MAC address " +
				"and DUID for identification. Requires IPv6 DHCP Server to have this option enabled as well to work properly.",
		},
		"binding_script": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A script that will be executed after binding is assigned or de-assigned. Internal `global` " +
				"variables that can be used in the script:\n    - bindingBound - set to `1` if bound, otherwise set to `0`\n" +
				"    - bindingServerName - dhcp server name\n    - bindingDUID - DUID\n    - bindingAddress - active " +
				"address\n    - bindingPrefix - active prefix.",
		},
		KeyComment: PropCommentRw,
		"dhcp_option": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Add additional DHCP options from option list.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		KeyDisabled: PropDisabledRw,
		"duid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "DUID value.",
		},
		KeyDynamic: PropDynamicRo,
		"ignore_ia_na_bindings": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Do not reply to DHCPv6 address requests and process only prefixes. Without this setting even " +
				"if server does not have address-pool configured, it has to respond to client that there is no address " +
				"available for the client. That can lead up to the situation when DHCPv6 client requests address and " +
				"prefix in a loop.",
		},
		"insert_queue_before": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify where to place dynamic simple queue entries for static DCHP leases with a " +
				"rate-limit parameter set.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"bottom", "first"}, false),
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The interface on which server will be running.",
		},
		KeyInvalid: PropInvalidRo,
		"lease_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time that a client may use the assigned address. The client will try to renew this address " +
				"after half of this time and will request a new address after the time limit expires.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyName: PropName("Reference name."),
		"parent_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A dynamically created queue for this lease will be configured as a child queue of the specified parent queue.",
		},
		"preference": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"prefix_pool": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IPv6 pool, from which to take IPv6 prefix for the clients.",
			AtLeastOneOf:     []string{"address_pool", "prefix_pool"},
		},
		"rapid_commit": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"route_distance": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Distance of the route.",
			ValidateFunc:     validation.IntBetween(1, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_radius": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to use RADIUS server.",
		},
		"use_reconfigure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Allow the server to send Reconfigure messages to clients.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
