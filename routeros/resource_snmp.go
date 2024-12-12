package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    "contact": "",
    "enabled": "false",
    "engine-id": "80003a8c04",
    "engine-id-suffix": "",
    "location": "",
    "src-address": "::",
    "trap-community": "public",
    "trap-generators": "temp-exception",
    "trap-target": "",
    "trap-version": "1",
    "vrf": "main"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/SNMP
func ResourceSNMP() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/snmp"),
		MetaId:           PropId(Id),

		"contact": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Contact information.",
		},
		KeyEnabled: PropEnabled("Used to disable/enable SNMP service"),
		"engine_id": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "For SNMP v3, used as part of identifier. You can configure suffix part of engine id " +
				"using this argument. If SNMP client is not  capable to detect set engine-id value then " +
				"this prefix hex have to be  used 0x80003a8c04",
		},
		"engine_id_suffix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Unique identifier for an SNMPv3 engine by configuring the suffix of the engine ID.",
		},
		"location": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Location information.",
		},
		"trap_community": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
			Description: "Which communities configured in community menu to use when sending out the trap. " +
				"This name must be present in the community list.",
		},
		"trap_generators": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "What action will generate traps:\n  * interfaces - interface changes;\n  * start-trap - snmp " +
				"server starting on the router.",
			ValidateFunc: validation.StringInSlice([]string{"interfaces", "start-trap", "temp-exception"}, false),
		},
		"trap_interfaces": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of interfaces that traps are going to be sent out.",
		},
		"trap_target": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "IP (IPv4 or IPv6) addresses of SNMP data collectors that have to receive the trap.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
		},
		"trap_version": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Version of SNMP protocol to use for trap.",
			ValidateFunc: validation.IntBetween(1, 3),
		},
		"src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Force the router to always use the same IP source address for all of the SNMP messages.",
			ValidateFunc: validation.IsIPAddress,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if (old == "" && new == "::") || (old == "::" && new == "") {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				return false
			},
		},
		KeyVrf: PropVrfRw,
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
