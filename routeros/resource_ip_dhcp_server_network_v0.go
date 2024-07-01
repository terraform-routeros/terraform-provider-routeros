package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServerNetwork https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Server#Networks
func ResourceDhcpServerNetworkV0() *schema.Resource {
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
			Type:     schema.TypeString,
			Optional: true,
			Description: "A comma-separated list of IP addresses for one or more CAPsMAN system managers. " +
				"DHCP Option 138 (capwap) will be used.",
		},
		KeyComment: PropCommentRw,
		"dhcp_option": {
			Type:        schema.TypeString,
			Optional:    true,
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
			Type:     schema.TypeString,
			Optional: true,
			Description: "the DHCP client will use these as the default DNS servers. Two comma-separated DNS servers " +
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
		"ntp_server": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The DHCP client will use these as the default NTP servers. Two comma-separated NTP servers " +
				"can be specified to be used by the DHCP client as primary and secondary NTP servers",
			ValidateFunc: validation.IsIPv4Address,
		},
		"wins_server": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The Windows DHCP client will use these as the default WINS servers. Two comma-separated " +
				"WINS servers can be specified to be used by the DHCP client as primary and secondary WINS servers",
			ValidateFunc: validation.IsIPv4Address,
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
