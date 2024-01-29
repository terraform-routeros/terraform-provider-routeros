package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInterfaceIPIPV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/IPIP"),
			MetaId:           PropId(Name),

			KeyActualMtu:     PropActualMtuRo,
			KeyAllowFastPath: PropAllowFastPathRw,
			KeyClampTcpMss:   PropClampTcpMssRw,
			KeyComment:       PropCommentRw,
			KeyDisabled:      PropDisabledRw,
			KeyDontFragment:  PropDontFragmentRw,
			KeyDscp:          PropDscpRw,
			KeyIpsecSecret:   PropIpsecSecretRw,
			KeyKeepalive:     PropKeepaliveRw,
			KeyL2Mtu:         PropL2MtuRo,
			KeyLocalAddress:  PropLocalAddressRw,
			KeyMtu:           PropMtuRw(),
			KeyName:          PropNameForceNewRw,
			KeyRemoteAddress: PropRemoteAddressRw,
			KeyRunning:       PropRunningRo,
		},
	}
}
