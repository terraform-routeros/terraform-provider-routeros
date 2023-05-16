package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 {
    "allow-fast-path": "true",
    "bridge-fast-forward-bytes": "0",
    "bridge-fast-forward-packets": "0",
    "bridge-fast-path-active": "true",
    "bridge-fast-path-bytes": "0",
    "bridge-fast-path-packets": "0",
    "use-ip-firewall": "false",
    "use-ip-firewall-for-pppoe": "false",
    "use-ip-firewall-for-vlan": "false"
  }
*/

// https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge#Bridge_Settings
func ResourceInterfaceBridgeSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/settings"),
		MetaId:           PropId(Name),

		"use_ip_firewall": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Force bridged traffic to also be processed by prerouting, forward and postrouting sections " +
				"of IP routing ( Packet Flow). This does not apply to routed traffic. This property is required in " +
				"case you want to assign Simple Queues or global Queue Tree to traffic in a bridge. Property " +
				"use-ip-firewall-for-vlan is required in case bridge vlan-filtering is used.",
		},
		"use_ip_firewall_for_pppoe": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Send bridged un-encrypted PPPoE traffic to also be processed by IP/Firewall. This " +
				"property only has effect when use-ip-firewall is set to yes. This property is required " +
				"in case you want to assign Simple Queues or global Queue Tree to PPPoE traffic in a " +
				"bridge.",
		},
		"use_ip_firewall_for_vlan": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Send bridged VLAN traffic to also be processed by IP/Firewall. This property only has " +
				"effect when use-ip-firewall is set to yes. This property is required in case you want " +
				"to assign Simple Queues or global Queue Tree to VLAN traffic in a bridge.",
		},
		"allow_fast_path": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to enable a bridge FastPath globally.",
		},
		"bridge_fast_path_active": {
			Type:     schema.TypeBool,
			Computed: true,
			Description: "Shows whether a bridge FastPath is active globally, FastPatch status per bridge " +
				"interface is not displayed.",
		},
		"bridge_fast_path_packets": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Shows packet count forwarded by Bridge FastPath.",
		},
		"bridge_fast_path_bytes": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Shows byte count forwarded by Bridge Fast Path.",
		},
		"bridge_fast_forward_packets": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Shows packet count forwarded by Bridge Fast Forward.",
		},
		"bridge_fast_forward_bytes": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Shows byte count forwarded by Bridge Fast Forward.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
