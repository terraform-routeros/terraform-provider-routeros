package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "action": "none",
    "common-name-regexp": "",
    "disabled": "false",
    "hw-supported-modes": "",
    "identity-regexp": "",
    "ip-address-ranges": "",
    "master-configuration": "cfg1",
    "name-format": "cap",
    "name-prefix": "",
    "radio-mac": "00:00:00:00:00:00",
    "slave-configurations": ""
  }
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManProvisioningV0() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/provisioning"),
		MetaId:           PropId(Id),

		"action": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Provisioning action.",
			ValidateFunc: validation.StringInSlice([]string{"create-disabled", "create-enabled",
				"create-dynamic-enabled", "none"}, false),
		},
		KeyComment: PropCommentRw,
		"common_name_regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Regular expression to match radios by common name. Each CAP's common name identifier can be " +
				`found under "/caps-man radio" as value "REMOTE-CAP-NAME"`,
		},
		KeyDisabled: PropDisabledRw,
		"hw_supported_modes": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Match radios by supported wireless modes.",
			ValidateFunc: validation.StringInSlice([]string{"a", "a-turbo", "ac", "an", "b", "g", "g-turbo", "gn"}, false),
		},
		"identity_regexp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Regular expression to match radios by router identity.",
		},
		"ip_address_ranges": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Match CAPs with IPs within configured address range.",
		},
		"master_configuration": {
			Type:     schema.TypeString,
			Required: true,
			Description: "If action specifies to create interfaces, then a new master interface with its configuration " +
				"set to this configuration profile will be created",
		},
		"name_format": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "cap",
			Description:  "Specify the syntax of the CAP interface name creation.",
			ValidateFunc: validation.StringInSlice([]string{"cap", "identity", "prefix", "prefix-identity"}, false),
		},
		"name_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name prefix which can be used in the name-format for creating the CAP interface names.",
		},
		"radio_mac": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "00:00:00:00:00:00",
			Description:  "MAC address of radio to be matched, empty MAC (00:00:00:00:00:00) means match all MAC addresses.",
			ValidateFunc: ValidationMacAddress,
		},
		"slave_configurations": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If action specifies to create interfaces, then a new slave interface for each configuration " +
				"profile in this list is created.",
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
