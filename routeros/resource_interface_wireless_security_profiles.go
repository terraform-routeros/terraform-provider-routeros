package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*0",
    "authentication-types": "wpa2-psk",
    "comment": "defconf",
    "default": "true",
    "disable-pmkid": "true",
    "eap-methods": "passthrough",
    "group-ciphers": "aes-ccm",
    "group-key-update": "5m",
    "interim-update": "0s",
    "management-protection": "disabled",
    "management-protection-key": "",
    "mode": "dynamic-keys",
    "mschapv2-password": "",
    "mschapv2-username": "",
    "name": "default",
    "radius-called-format": "mac:ssid",
    "radius-eap-accounting": "false",
    "radius-mac-accounting": "false",
    "radius-mac-authentication": "false",
    "radius-mac-caching": "disabled",
    "radius-mac-format": "XX:XX:XX:XX:XX:XX",
    "radius-mac-mode": "as-username",
    "static-algo-0": "none",
    "static-algo-1": "none",
    "static-algo-2": "none",
    "static-algo-3": "none",
    "static-key-0": "",
    "static-key-1": "",
    "static-key-2": "",
    "static-key-3": "",
    "static-sta-private-algo": "none",
    "static-sta-private-key": "",
    "static-transmit-key": "key-0",
    "supplicant-identity": "MikroTik",
    "tls-certificate": "none",
    "tls-mode": "no-certificates",
    "unicast-ciphers": "aes-ccm",
    "wpa-pre-shared-key": "",
    "wpa2-pre-shared-key": "XXX"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Wireless+Interface#WirelessInterface-SecurityProfiles
func ResourceInterfaceWirelessSecurityProfiles() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless/security-profiles"),
		MetaId:           PropId(Id),

		"authentication_types": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Set of supported authentication types, multiple values can be selected. Access Point will " +
				"advertise supported authentication types, and client will connect to Access Point only if it supports " +
				"any of the advertised authentication types.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wpa-psk", "wpa2-psk", "wpa-eap", "wpa2-eap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		KeyDefault: PropDefaultRo,
		"disable_pmkid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to include `PMKID` into the `EAPOL` frame sent out by the Access Point. Disabling PMKID " +
				"can cause compatibility issues with devices that use the PMKID to connect to an Access Point. `yes` - removes " +
				"PMKID from EAPOL frames (improves security, reduces compatibility). `no` - includes PMKID into EAPOL frames " +
				"(reduces security, improves compatibility).This property only has effect on Access Points.",
		},
		"eap_methods": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allowed types of authentication methods, multiple values can be selected. This property only " +
				"has effect on Access Points. `eap-tls` - Use built-in EAP TLS authentication. Both client and server certificates " +
				"are supported. See description of tls-mode and tls-certificate properties. `eap-ttls-mschapv2` - Use EAP-TTLS " +
				"with MS-CHAPv2 authentication. `passthrough` - Access Point will relay authentication process to the RADIUS " +
				"server. `peap` - Use Protected EAP authentication.",
			ValidateFunc:     validation.StringInSlice([]string{"eap-tls", "eap-ttls-mschapv2", "passthrough", "peap"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"group_ciphers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point advertises one of these ciphers, multiple values can be selected. Access Point " +
				"uses it to encrypt all broadcast and multicast frames. Client attempts connection only to Access Points " +
				"that use one of the specified group ciphers. `tkip` - Temporal Key Integrity Protocol - encryption protocol, " +
				"compatible with legacy WEP equipment, but enhanced to correct some of the WEP flaws. `aes-ccm` - more secure " +
				"WPA encryption protocol, based on the reliable AES (Advanced Encryption Standard). Networks free of " +
				"WEP legacy should use only this cipher.",
			ValidateFunc:     validation.StringInSlice([]string{"tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"group_key_update": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Controls how often Access Point updates the group key. This key is used to encrypt all broadcast " +
				"and multicast frames. property only has effect for Access Points.",
			DiffSuppressFunc: TimeEquall,
		},
		"interim_update": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When RADIUS accounting is used, Access Point periodically sends accounting information updates " +
				"to the RADIUS server. This property specifies default update interval that can be overridden by the " +
				"RADIUS server using Acct-Interim-Interval attribute.",
			DiffSuppressFunc: TimeEquall,
		},
		"management_protection": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Management frame protection. Used for: Deauthentication attack prevention, MAC address cloning " +
				"issue. Possible values are: `disabled` - management protection is disabled (default), `allowed` - use " +
				"management protection if supported by remote party (for AP - allow both, non-management protection " +
				"and management protection clients, for client - connect both to APs with and without management " +
				"protection), `required` - establish association only with remote devices that support management " +
				"protection (for AP - accept only clients that support management protection, for client - connect " +
				"only to APs that support management protection).",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "allowed", "required"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"management_protection_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Management protection shared secret. When interface is in AP mode, default management " +
				"protection key (configured in security-profile) can be overridden by key specified in access-list or " +
				"RADIUS attribute.",
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Encryption mode for the security profile. `none` - Encryption is not used. Encrypted frames are " +
				"not accepted. `static-keys-required` - WEP mode. Do not accept and do not send unencrypted frames. Station " +
				"in static-keys-required mode will not connect to an Access Point in static-keys-optional mode. `static-keys-optional` - " +
				"WEP mode. Support encryption and decryption, but allow also to receive and send unencrypted frames. " +
				"Device will send unencrypted frames if encryption algorithm is specified as none. Station in static-keys-optional mode " +
				"will not connect to an Access Point in `static-keys-required` mode. See also: static-sta-private-algo, " +
				"static-transmit-key. `dynamic-keys` - WPA mode.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "static-keys-optional", "static-keys-required", "dynamic-keys"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mschapv2_password": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Password to use for authentication when `eap-ttls-mschapv2` or `peap` authentication method is " +
				"being used. This property only has effect on Stations.",
		},
		"mschapv2_username": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Username to use for authentication when `eap-ttls-mschapv2` or `peap` authentication method is " +
				"being used. This property only has effect on Stations.",
		},
		KeyName: PropName("Name of the security profile."),
		"radius_called_format": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "mac | mac:ssid | ssid",
			ValidateFunc:     validation.StringInSlice([]string{"mac", "mac:ssid", "ssid"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius_eap_accounting": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"radius_mac_accounting": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"radius_mac_authentication": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "This property affects the way how Access Point processes clients that are not found in the Access " +
				"List.no - allow or reject client authentication based on the value of default-authentication property " +
				"of the Wireless interface.yes - Query RADIUS server using MAC address of client as user name. With this " +
				"setting the value of default-authentication has no effect.",
		},
		"radius_mac_caching": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If this value is set to time interval, the Access Point will cache RADIUS MAC authentication " +
				"responses for specified time, and will not contact RADIUS server if matching cache entry already exists. " +
				"Value disabled will disable cache, Access Point will always contact RADIUS server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius_mac_format": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Controls how MAC address of the client is encoded by Access Point in the User-Name attribute " +
				"of the MAC authentication and MAC accounting RADIUS requests.",
			ValidateFunc: validation.StringInSlice([]string{"XX:XX:XX:XX:XX:XX", "XXXX:XXXX:XXXX", "XXXXXX:XXXXXX",
				"XX-XX-XX-XX-XX-XX", "XXXXXX-XXXXXX", "XXXXXXXXXXXX", "XX XX XX XX XX XX"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius_mac_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "By default Access Point uses an empty password, when sending Access-Request during MAC authentication. " +
				"When this property is set to `as-username-and-password`, Access Point will use the same value for User-Password " +
				"attribute as for the User-Name attribute.",
			ValidateFunc:     validation.StringInSlice([]string{"as-username", "as-username-and-password"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_algo_0": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Encryption algorithm to use with the corresponding key.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_algo_1": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Encryption algorithm to use with the corresponding key.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_algo_2": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Encryption algorithm to use with the corresponding key.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_algo_3": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Encryption algorithm to use with the corresponding key.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_key_0": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.",
		},
		"static_key_1": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.",
		},
		"static_key_2": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.",
		},
		"static_key_3": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.",
		},
		"static_sta_private_algo": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Encryption algorithm to use with station private key. Value none disables use of the private " +
				"key. This property is only used on Stations. Access Point has to get corresponding value either from private-algo property, " +
				"or from Mikrotik-Wireless-Enc-Algo attribute. Station private key replaces key 0 for unicast frames. " +
				"Station will not use private key to decrypt broadcast frames.",
			ValidateFunc:     validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"static_sta_private_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Length of key must be appropriate for selected algorithm, see the Statically configured WEP " +
				"keys section. This property is used only on Stations. Access Point uses corresponding key either from " +
				"private-key property, or from Mikrotik-Wireless-Enc-Key attribute.",
		},
		"static_transmit_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point will use the specified key to encrypt frames for clients that do not use private " +
				"key. Access Point will also use this key to encrypt broadcast and multicast frames. Client will use " +
				"the specified key to encrypt frames if static-sta-private-algo is set to none. If corresponding static-algo-N property " +
				"has value set to none, then frame will be sent unencrypted (when mode is set to static-keys-optional) " +
				"or will not be sent at all (when mode is set to static-keys-required).",
			ValidateFunc:     validation.StringInSlice([]string{"key-0", "key-1", "key-2", "key-3"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"supplicant_identity": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "EAP identity that is sent by client at the beginning of EAP authentication. This value is " +
				"used as a value for User-Name attribute in RADIUS messages sent by RADIUS EAP accounting and RADIUS " +
				"EAP pass-through authentication.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tls_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point always needs a certificate when configured when tls-mode is set to verify-certificate, " +
				"or is set to dont-verify-certificate. Client needs a certificate only if Access Point is configured " +
				"with tls-mode set to verify-certificate. In this case client needs a valid certificate that is signed " +
				"by a CA known to the Access Point. This property only has effect when tls-mode is not set to " +
				"no-certificates and eap-methods contains eap-tls.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tls_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This property has effect only when eap-methods contains eap-tls. `verify-certificate` - Require " +
				"remote device to have valid certificate. Check that it is signed by known certificate authority. No " +
				"additional identity verification is done. Certificate may include information about time period during " +
				"which it is valid. If router has incorrect time and date, it may reject valid certificate because router's " +
				"clock is outside that period. See also the Certificates configuration. `dont-verify-certificate` - Do not " +
				"check certificate of the remote device. Access Point will not require client to provide certificate. " +
				"`no-certificates` - Do not use certificates. TLS session is established using 2048 bit anonymous " +
				"Diffie-Hellman key exchange. `verify-certificate-with-crl` - Same as verify-certificate but also " +
				"checks if the certificate is valid by checking the Certificate Revocation List.",
			ValidateFunc: validation.StringInSlice([]string{"verify-certificate", "dont-verify-certificate",
				"no-certificates", "verify-certificate-with-crl"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"unicast_ciphers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Access Point advertises that it supports specified ciphers, multiple values can be selected. " +
				"Client attempts connection only to Access Points that supports at least one of the specified ciphers. " +
				"One of the ciphers will be used to encrypt unicast frames that are sent between Access Point and Station.",
			ValidateFunc:     validation.StringInSlice([]string{"tkip", "aes-ccm"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wpa_pre_shared_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "`WPA` pre-shared key mode requires all devices in a BSS to have common secret key. Value of " +
				"this key can be an arbitrary text. Commonly referred to as the network password for WPA mode. property " +
				"only has effect when wpa-psk is added to authentication-types.",
			ValidateFunc: validation.StringLenBetween(8, 64),
		},
		"wpa2_pre_shared_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "`WPA2` pre-shared key mode requires all devices in a BSS to have common secret key. Value of " +
				"this key can be an arbitrary text. Commonly referred to as the network password for WPA2 mode. property " +
				"only has effect when wpa2-psk is added to authentication-types.",
			ValidateFunc: validation.StringLenBetween(8, 64),
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
