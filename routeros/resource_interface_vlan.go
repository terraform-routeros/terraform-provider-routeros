package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceInterfaceVlan https://wiki.mikrotik.com/wiki/Manual:Interface/VLAN
func ResourceInterfaceVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vlan"),
		MetaId:           PropId(Id),

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
		"mvrp": {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  false,
			Description: "Specifies whether this VLAN should declare its attributes through Multiple VLAN Registration Protocol (MVRP) as an applicant (available since RouterOS 7.15). " +
				"It can be used to register the VLAN with connected bridges that support MVRP. " +
				"This property only has an effect when use-service-tag is disabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName:    PropNameForceNewRw,
		KeyRunning: PropRunningRo,
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

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceInterfaceVlanV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
