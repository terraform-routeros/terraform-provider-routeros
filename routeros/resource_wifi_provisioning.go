package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "action": "create-enabled",
    "address-ranges": "192.168.88.1-192.168.88.100,192.168.100.1-192.168.100.100",
    "common-name-regexp": "test",
    "disabled": "false",
    "identity-regexp": "test",
    "master-configuration": "cfg1",
    "name-format": "cap1:",
    "radio-mac": "00:00:00:00:00:00",
    "slave-configurations": "cfg1",
    "supported-bands": "2ghz-n,2ghz-g"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-CAPsMANProvisioning
func ResourceWifiProvisioning() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/provisioning"),
		MetaId:           PropId(Id),

		"action": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Provisioning action.",
			ValidateFunc: validation.StringInSlice([]string{"create-disabled", "create-enabled",
				"create-dynamic-enabled", "none"}, false),
		},
		"address_ranges": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Match CAPs by IPs within configured address ranges.",
		},
		KeyComment: PropCommentRw,
		"common_name_regexp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Regular expression to match radios by common name.",
		},
		KeyDisabled: PropDisabledRw,
		"identity_regexp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Regular expression to match radios by router identity.",
		},
		"master_configuration": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If action specifies to create interfaces, then a new master interface with its configuration " +
				"set to this configuration profile will be created.",
		},
		"name_format": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specify the format of the CAP interface name creation.",
		},
		"radio_mac": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "MAC address of radio to be matched, empty MAC means match all MAC addresses. `00:00:00:00:00:00` is not considered empty MAC-address.",
			ValidateFunc:     ValidationMacAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"slave_configurations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "If action specifies to create interfaces, then a new slave interface for each configuration " +
				"profile in this list is created.",
		},
		"slave_name_format": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The name format of the slave CAP interfaces. This option is available in RouterOS starting from version 7.16.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"supported_bands": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"2ghz-ax", "2ghz-g", "2ghz-n", "5ghz-a", "5ghz-ac", "5ghz-ax", "5ghz-n"}, false),
			},
			Description: "Match CAPs by supported modes.",
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
