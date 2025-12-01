package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "arp": "disabled",
    "bridge": "bridge",
    "bridge-cost": "0",
    "bridge-horizon": "none",
    "client-to-client-forwarding": "true",
    "interface-list": "all",
    "l2mtu": "0",
    "local-forwarding": "true",
    "mtu": "32",
    "name": "datapath1",
    "vlan-id": "1",
    "vlan-mode": "no-tag"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManDatapath() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/datapath"),
		MetaId:           PropId(Id),

		"arp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ARP mode. See [docs](https://wiki.mikrotik.com/wiki/Manual:IP/ARP#ARP_Modes) for info.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled", "local-proxy-arp", "proxy-arp",
				"reply-only"}, false),
		},
		"bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Bridge to which particular interface should be automatically added as port. Required " +
				"only when local-forwarding is not used.",
		},
		"bridge_cost": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Bridge port cost to use when adding as bridge port.",
		},
		"bridge_horizon": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Bridge horizon to use when adding as bridge port.",
		},
		"client_to_client_forwarding": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Controls if client-to-client forwarding between wireless clients connected to interface " +
				"should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is " +
				"performed by CAPsMAN.",
		},
		KeyComment: PropCommentRw,
		"interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface list name.",
		},
		"l2mtu": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Layer2 MTU size.",
		},
		"local_forwarding": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Controls forwarding mode. If disabled, all L2 and L3 data will be forwarded to CAPsMAN, " +
				"and further forwarding decisions will be made only then. See [docs](https://wiki.mikrotik.com/wiki/Manual:CAPsMAN#Local_Forwarding_Mode) for info.",
		},
		KeyName: PropNameForceNewRw,
		"mtu": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "MTU size.",
		},
		"openflow_switch": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "OpenFlow switch to add interface to, as port when enabled.",
		},
		KeyVlanId: PropVlanIdRw("VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging.", false),
		"vlan_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "VLAN tagging mode specifies if VLAN tag should be assigned to interface (causes all received " +
				"data to get tagged with VLAN tag and allows interface to only send out data tagged with given tag)",
			ValidateFunc: validation.StringInSlice([]string{"no-tag", "use-service-tag", "use-tag"}, false),
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
				Type:    ResourceCapsManDatapathV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
