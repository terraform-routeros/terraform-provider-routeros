package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "authentication-types": "",
    "disable-pmkid": "false",
    "eap-methods": "eap-tls",
    "eap-radius-accounting": "false",
    "encryption": "",
    "group-encryption": "aes-ccm",
    "group-key-update": "1m",
    "name": "security1",
    "passphrase": "123123123",
    "tls-certificate": "none",
    "tls-mode": "verify-certificate"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManSecurity() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/security"),
		MetaId:           PropId(Id),

		"authentication_types": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Specify the type of Authentication from wpa-psk, wpa2-psk, wpa-eap or wpa2-eap.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"wpa-psk", "wpa2-psk", "wpa-eap", "wpa2-eap"},
					false, false),
			},
		},
		KeyComment: PropCommentRw,
		"disable_pmkid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to include PMKID into the EAPOL frame sent out by the Access Point. Disabling PMKID " +
				"can cause compatibility issues with devices that use the PMKID to connect to an Access Point.",
		},
		"eap_methods": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "eap-tls - Use built-in EAP TLS authentication; passthrough - Access point will relay " +
				"authentication process to the RADIUS server.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"eap-tls", "passthrough"}, false, false),
		},
		"eap_radius_accounting": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies if RADIUS traffic accounting should be used if RADIUS authentication gets done for " +
				"this client",
		},
		"encryption": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Set type of unicast encryption algorithm used.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"aes-ccm", "tkip"}, false, false),
			},
		},
		"group_encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point advertises one of these ciphers, multiple values can be selected. Access Point " +
				"uses it to encrypt all broadcast and multicast frames. Client attempts connection only to Access " +
				"Points that use one of the specified group ciphers.",
			ValidateFunc: validation.StringInSlice([]string{"aes-ccm", "tkip"}, false),
		},
		"group_key_update": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Controls how often Access Point updates the group key. This key is used to encrypt all " +
				"broadcast and multicast frames. property only has effect for Access Points. (30s..1h)",
			DiffSuppressFunc: TimeEqual,
		},
		KeyName: PropNameForceNewRw,
		// 802.11i specification:
		// A pass-phrase is a sequence of between 8 and 63 ASCII-encoded characters. The limit of 63 comes from the
		// desire to distinguish between a pass-phrase and a PSK displayed as 64 hexadecimal characters.
		"passphrase": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "WPA or WPA2 pre-shared key.",
			Sensitive:    true,
			ValidateFunc: validation.StringLenBetween(8, 63),
		},
		"tls_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point always needs a certificate when security.tls-mode is set to value other than " +
				"no-certificates.",
		},
		"tls_mode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "This property has effect only when security.eap-methods contains eap-tls.",
			ValidateFunc: validation.StringInSlice([]string{"verify-certificate", "dont-verify-certificate",
				"no-certificates", "verify-certificate-with-crl"}, false),
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceCapsManSecurityV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
