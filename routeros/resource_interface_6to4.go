package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*3",
    "actual-mtu": "1480",
    "clamp-tcp-mss": "true",
    "disabled": "false",
    "dont-fragment": "no",
    "dscp": "inherit",
    "local-address": "0.0.0.0",
    "mtu": "auto",
    "name": "6to4-tunnel1",
    "remote-address": "unspecified",
    "running": "true"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/135004174/6to4
func ResourceInterface6to4() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/6to4"),
		MetaId:           PropId(Id),

		KeyActualMtu:     PropActualMtuRo,
		KeyClampTcpMss:   PropClampTcpMssRw,
		KeyComment:       PropCommentRw,
		KeyDisabled:      PropDisabledRw,
		KeyDontFragment:  PropDontFragmentRw,
		KeyDscp:          PropDscpRw,
		KeyIpsecSecret:   PropIpsecSecretRw,
		KeyKeepalive:     PropKeepaliveRw,
		KeyLocalAddress:  PropLocalAddressRw,
		KeyMtu:           PropMtuRw(),
		KeyName:          PropName("Interface name."),
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
