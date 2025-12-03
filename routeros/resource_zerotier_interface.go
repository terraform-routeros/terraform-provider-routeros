package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "allow-default": "false",
    "allow-global": "false",
    "allow-managed": "false",
    "arp-timeout": "auto",
    "bridge": "true",
    "dhcp": "false",
    "disabled": "false",
    "instance": "zt1",
    "mac-address": "00:00:00:00:00:00",
    "mtu": "2800",
    "name": "zerotier1",
    "network": "a00000000aa00a00",
    "network-name": "something",
    "running": "true",
    "status": "OK",
    "type": "PRIVATE"
}
*/

// https://help.mikrotik.com/docs/display/ROS/ZeroTier#ZeroTier-Parameters
func ResourceZerotierInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/zerotier/interface"),
		MetaId:           PropId(Id),

		"allow_default": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option to override the default route.",
		},
		"allow_global": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option to allow overlapping public IP space by the ZeroTier routes. .",
		},
		"allow_managed": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option to allow assignment of managed IPs.",
		},
		KeyArpTimeout: PropArpTimeoutRw,
		"bridge": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the ZeroTier interface is bridged.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"disable_running_check": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to force the `running` property to true.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dhcp": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the ZeroTier interface obtained an IP address.",
		},
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier instance name.",
		},
		KeyMacAddress: PropMacAddressRo,
		KeyMtu:        PropL2MtuRo,
		KeyName:       PropName("Name of the ZeroTier interface."),
		"network": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier network identifier.",
		},
		"network_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ZeroTier network name.",
		},
		KeyRunning: PropRunningRo,
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of the ZeroTier connection.",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ZeroTier network type.",
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
