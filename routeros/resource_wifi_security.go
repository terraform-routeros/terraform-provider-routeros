package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "authentication-types": "wpa2-psk,wpa3-psk",
    "connect-group": "something",
    "connect-priority": "0",
    "dh-groups": "19,20,21",
    "disable-pmkid": "false",
    "disabled": "false",
    "eap-accounting": "true",
    "eap-anonymous-identity": "anonymous",
    "eap-certificate-mode": "verify-certificate",
    "eap-methods": "peap,tls",
    "eap-password": "",
    "eap-tls-certificate": "ca",
    "eap-username": "",
    "encryption": "tkip",
    "ft": "true",
    "ft-mobility-domain": "0x1",
    "ft-nas-identifier": "ssid",
    "ft-over-ds": "true",
    "ft-preserve-vlanid ": "true",
    "ft-r0-key-lifetime": "10m",
    "ft-reassociation-deadline": "10s",
    "group-encryption": "tkip",
    "group-key-update": "10m",
    "management-encryption": "cmac",
    "management-protection": "disabled",
    "name": "sec1",
    "passphrase": "passphrase",
    "sae-anti-clogging-threshold": "0",
    "sae-max-failure-rate": "disabled",
    "sae-pwe": "hunting-and-pecking",
    "wps": "disable"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-SecurityProperties
func ResourceWifiSecurity() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/security"),
		MetaId:           PropId(Id),

		"authentication_types": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wpa-psk", "wpa2-psk", "wpa-eap", "wpa2-eap", "wpa3-psk", "owe", "wpa3-eap", "wpa3-eap-192"}, false),
			},
			Description: "Authentication types to enable on the interface.",
		},
		KeyComment: PropCommentRw,
		"connect_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "APs within the same connect group do not allow more than 1 client device with the same MAC address.",
		},
		"connect_priority": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An option to determine how a connection is handled if the MAC address of the client device is the same as that of another active connection to another AP.",
		},
		"dh_groups": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{19, 20, 21}),
			},
			Description: "Identifiers of elliptic curve cryptography groups to use in SAE (WPA3) authentication.",
		},
		"disable_pmkid": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to disable inclusion of a PMKID in EAPOL frames.",
		},
		KeyDisabled: PropDisabledRw,
		"eap_accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to send accounting information to RADIUS server for EAP-authenticated peers.",
		},
		"eap_anonymous_identity": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An option to specify anonymous identity for EAP outer authentication.",
		},
		"eap_certificate_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "A policy for handling the TLS certificate of the RADIUS server.",
			ValidateFunc: validation.StringInSlice([]string{"dont-verify-certificate", "no-certificates", "verify-certificate", "verify-certificate-with-crl"}, false),
		},
		"eap_methods": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"peap", "tls", "ttls"}, false),
			},
			Description: "A set of EAP methods to consider for authentication.",
		},
		"eap_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Password to use when the chosen EAP method requires one.",
		},
		"eap_tls_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name or id of a certificate in the device's certificate store to use when the chosen EAP authentication method requires one.",
		},
		"eap_username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Username to use when the chosen EAP method requires one. ",
		},
		"encryption": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ccmp", "ccmp-256", "gcmp", "gcmp-256", "tkip"}, false),
			},
			Description: "A list of ciphers to support for encrypting unicast traffic.",
		},
		"ft": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable 802.11r fast BSS transitions (roaming).",
		},
		"ft_mobility_domain": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The fast BSS transition mobility domain ID.",
			ValidateFunc: Validation64k,
		},
		"ft_nas_identifier": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Fast BSS transition PMK-R0 key holder identifier.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[0-9a-zA-Z]{2,96}$`),
				"Must be a string of 2 - 96 hex characters."),
		},
		"ft_over_ds": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to enable fast BSS transitions over DS (distributed system).",
		},
		"ft_preserve_vlanid": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "An option to preserve VLAN ID when roaming.",
		},
		"ft_r0_key_lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The lifetime of the fast BSS transition PMK-R0 encryption key.",
			DiffSuppressFunc: TimeEquall,
		},
		"ft_reassociation_deadline": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fast BSS transition reassociation deadline.",
			DiffSuppressFunc: TimeEquall,
		},
		"group_encryption": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "A cipher to use for encrypting multicast traffic.",
			ValidateFunc: validation.StringInSlice([]string{"ccmp", "ccmp-256", "gcmp", "gcmp-256", "tkip"}, false),
		},
		"group_key_update": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval at which the group temporal key (key for encrypting broadcast traffic) is renewed.",
			DiffSuppressFunc: TimeEquall,
		},
		"management_encryption": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "A cipher to use for encrypting protected management frames.",
			ValidateFunc: validation.StringInSlice([]string{"cmac", "cmac-256", "gmac", "gmac-256"}, false),
		},
		"management_protection": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to enable 802.11w management frame protection.",
			ValidateFunc: validation.StringInSlice([]string{"allowed", "disabled", "required"}, false),
		},
		"multi_passphrase_group": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of `/interface/wifi/security/multi-passphrase/` group that will be used. Only a " +
				"single group can be defined under the security profile.",
		},
		"owe_transition_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name or internal ID of an interface which MAC address and SSID to advertise as the matching AP when running in OWE transition mode.",
		},
		KeyName: PropName("Name of the security profile."),
		"passphrase": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Passphrase to use for PSK authentication types.",
		},
		"sae_anti_clogging_threshold": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A parameter to mitigate DoS attacks by specifying a threshold of in-progress SAE authentications.",
		},
		"sae_max_failure_rate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Rate of failed SAE (WPA3) associations per minute, at which the AP will stop processing new association requests.",
		},
		"sae_pwe": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Methods to support for deriving SAE password element.",
			ValidateFunc: validation.StringInSlice([]string{"both", "hash-to-element", "hunting-and-pecking"}, false),
		},
		"wps": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to enable WPS (Wi-Fi Protected Setup).",
			ValidateFunc: validation.StringInSlice([]string{"disable", "push-button"}, false),
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
