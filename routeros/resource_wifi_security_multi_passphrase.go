package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "expired": "false",
    "group": "123",
    "passphrase": "12345678"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/224559120/WiFi#WiFi-Securitymulti-passphraseproperties
func ResourceWifiSecurityMultiPassphrase() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/security/multi-passphrase"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"expires": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The expiration date and time for passphrase specified in this entry, doesn't affect the whole " +
				"group. Once the date is reached, existing clients using this passphrase will be disconnected, and new " +
				"clients will not be able to connect using it. If not set, passphrase can be used indefinetly.",
		},
		"group": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Assigning the group to a security profile or an access list, will enable use of all passphrases " +
				"defined under it.",
		},
		"isolation": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Determines whether the client device using this passphrase is isolated from other clients " +
				"on AP. Traffic from an isolated client will not be forwarded to other clients and unicast traffic from " +
				"a non-isolated client will not be forwarded to an isolated one.",
		},
		"passphrase": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "The passphrase to use for PSK authentication types. Multiple users can use the same passphrase. " +
				"Not compatible with WPA3-PSK.",
			ValidateFunc: validation.StringLenBetween(8, 64),
		},
		"vlan_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Vlan-id that will be assigned to clients using this passphrase Only supported on wifi-qcom " +
				"interfaces, if wifi-qcom-ac AP has a client that uses a passphrase that has vlan-id associated with " +
				"it, the client will not be able to join.",
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.17beta1.</span>*`,
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
