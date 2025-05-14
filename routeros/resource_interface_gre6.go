package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceInterfaceGre6
func ResourceInterfaceGre6() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/gre6"),
		MetaId:           PropId(Id),

		KeyActualMtu:     PropActualMtuRo,
		KeyClampTcpMss:   PropClampTcpMssRw,
		KeyComment:       PropCommentRw,
		KeyDisabled:      PropDisabledRw,
		KeyDscp:          PropDscpRw,
		KeyIpsecSecret:   PropIpsecSecretRw,
		KeyKeepalive:     PropKeepaliveRw,
		KeyL2Mtu:         PropL2MtuRo,
		KeyLocalAddress:  PropLocalAddressRw,
		KeyMtu:           PropMtuRw(),
		KeyName:          PropNameForceNewRw,
		KeyRemoteAddress: PropRemoteAddressRw,
		KeyRunning:       PropRunningRo,
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
