package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceMacVlan https://help.mikrotik.com/docs/display/ROS/MACVLAN
func ResourceInterfaceMacVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/macvlan"),
		MetaId:           PropId(Id),

		KeyArp:                     PropArpRw,
		KeyArpTimeout:              PropArpTimeoutRw,
		KeyComment:                 PropCommentRw,
		KeyDisabled:                PropDisabledRw,
		KeyInterface:               PropInterfaceRw,
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "bridge",
			Description: "Sets MACVLAN interface mode:\n  *	private - does not allow communication between MACVLAN " +
				"instances on the same parent interface.\n  * bridge - allows communication between MACVLAN instances on " +
				"the same parent interface.",
			ValidateFunc: validation.StringInSlice([]string{"private", "bridge"}, true),
		},
		KeyMacAddress: PropMacAddressRw(
			`Static MAC address of the interface. A randomly generated MAC address will be assigned when not specified.`,
			false,
		),
		KeyL2Mtu: PropL2MtuRo,
		KeyName:  PropNameForceNewRw,
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
