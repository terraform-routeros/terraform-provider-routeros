package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceEoipV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/eoip"),
			MetaId:           PropId(Name),

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
		},
	}
}
