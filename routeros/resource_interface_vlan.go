package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceInterfaceVlan https://wiki.mikrotik.com/wiki/Manual:Interface/VLAN
func ResourceInterfaceVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:           PropResourcePath("/interface/vlan"),
		MetaId:                     PropId(Name),

		KeyArp:                     PropArpRw,
		KeyArpTimeout:              PropArpTimeoutRw,
		KeyComment:                 PropCommentRw,
		KeyDisabled:                PropDisabledRw,
		KeyInterface:               PropInterfaceRw,
		KeyL2Mtu:                   PropL2MtuRo,
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		KeyLoopProtectStatus:       PropLoopProtectStatusRo,
		KeyMacAddress:              PropMacAddressRo,
		KeyMtu:                     PropMtuRw(),
		KeyName:                    PropNameForceNewRw,
		KeyRunning:                 PropRunningRo,
		"use_service_tag": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vlan_id": {
			Type:     schema.TypeInt,
			Required: true,
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
