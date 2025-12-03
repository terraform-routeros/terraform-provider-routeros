package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/EoIP
func ResourceInterfaceEoip() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/eoip"),
		MetaId:           PropId(Id),

		KeyActualMtu:               PropActualMtuRo,
		KeyArp:                     PropArpRw,
		KeyArpTimeout:              PropArpTimeoutRw,
		KeyAllowFastPath:           PropAllowFastPathRw,
		KeyClampTcpMss:             PropClampTcpMssRw,
		KeyComment:                 PropCommentRw,
		KeyDisabled:                PropDisabledRw,
		KeyDontFragment:            PropDontFragmentRw,
		KeyDscp:                    PropDscpRw,
		KeyIpsecSecret:             PropIpsecSecretRw,
		KeyKeepalive:               PropKeepaliveRw,
		KeyL2Mtu:                   PropL2MtuRo,
		KeyLocalAddress:            PropLocalAddressRw,
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		KeyLoopProtectStatus:       PropLoopProtectStatusRo,
		KeyMacAddress:              PropMacAddressRo,
		KeyMtu:                     PropMtuRw(),
		KeyName:                    PropNameForceNewRw,
		KeyRemoteAddress:           PropRemoteAddressRw,
		KeyRunning:                 PropRunningRo,
		"tunnel_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Unique tunnel identifier, which must match the other side of the tunnel.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultValidateCreate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		ReadContext: DefaultRead(resSchema),
		UpdateContext: DefaultValidateUpdate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		DeleteContext: DefaultDelete(resSchema),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceInterfaceEoipV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
