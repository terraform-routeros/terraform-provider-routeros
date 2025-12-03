package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServerNetwork https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP#DHCP-Network
func ResourceDhcpServerNetwork() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/network"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The network DHCP server(s) will lease addresses from.",
		},
		"boot_file_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Boot filename.",
		},
		"caps_manager": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A list of IP addresses for one or more CAPsMAN system managers. " +
				"DHCP Option 138 (capwap) will be used.",
		},
		KeyComment: PropCommentRw,
		"dhcp_option": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Add additional DHCP options from the option list.",
		},
		"dhcp_option_set": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Add an additional set of DHCP options.",
		},
		"dns_none": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If set, then DHCP Server will not pass dynamic DNS servers configured on the router to the " +
				"DHCP clients if no DNS Server in DNS-server is set.",
		},
		"dns_server": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The DHCP client will use these as the default DNS servers. Two DNS servers " +
				"can be specified to be used by the DHCP client as primary and secondary DNS servers.",
		},
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The DHCP client will use this as the 'DNS domain' setting for the network adapter.",
		},
		KeyDynamic: PropDynamicRo,
		"gateway": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0.0.0.0",
			Description:  "The default gateway to be used by DHCP Client.",
			ValidateFunc: validation.IsIPv4Address,
		},
		"netmask": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
			Description: "The actual network mask is to be used by the DHCP client. If set to '0' - netmask from " +
				"network address will be used.",
			ValidateFunc: validation.IntBetween(0, 32),
		},
		"next_server": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The IP address of the next server to use in bootstrap.",
			ValidateFunc: validation.IsIPv4Address,
		},
		"ntp_none": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If set, then DHCP Server will not pass NTP servers configured on the router to the DHCP clients.",
		},
		"ntp_server": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			Description: "The DHCP client will use these as the default NTP servers. Two NTP servers " +
				"can be specified to be used by the DHCP client as primary and secondary NTP servers",
		},
		"wins_server": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
			},
			Description: "The Windows DHCP client will use these as the default WINS servers. Two WINS servers " +
				"can be specified to be used by the DHCP client as primary and secondary WINS servers",
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

		Schema:        resSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceDhcpServerNetworkV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationScalarToList("caps_manager", "dhcp_option", "dns_server", "ntp_server", "wins_server"),
				Version: 0,
			},
		},
	}
}
