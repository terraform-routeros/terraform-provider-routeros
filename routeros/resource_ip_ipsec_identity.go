package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*A",
    "auth-method": "pre-shared-key",
    "disabled": "false",
    "dynamic": "false",
    "generate-policy": "no",
    "peer": "peer1",
    "secret": "secret!!!"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/IPsec#IPsec-Identities
func ResourceIpIpsecIdentity() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/identity"),
		MetaId:           PropId(Id),

		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Authentication method:\n  * digital-signature - authenticate using a pair of RSA certificates;\n  * eap " +
				"- IKEv2 EAP authentication for initiator (peer with a netmask of `/32`). Must be used together with eap-methods;\n  * eap-radius " +
				"- IKEv2 EAP RADIUS passthrough authentication for the responder (RFC 3579). A server certificate in " +
				"this case is required. If a server certificate is not specified then only clients supporting EAP-only " +
				"(RFC 5998) will be able to connect. Note that the EAP method should be compatible with EAP-only;\n  * pre-shared-key " +
				"- authenticate by a password (pre-shared secret) string shared between the peers (not recommended since " +
				"an offline attack on the pre-shared key is possible);\n  * rsa-key - authenticate using an RSA key imported " +
				"in keys menu. Only supported in IKEv1;\n  * pre-shared-key-xauth - authenticate by a password (pre-shared " +
				"secret) string shared between the peers + XAuth username and password. Only supported in IKEv1;\n  * rsa-signature-hybrid " +
				"- responder certificate authentication with initiator XAuth. Only supported in IKEv1.",
			ValidateFunc: validation.StringInSlice([]string{"digital-signature", "eap", "eap-radius", "pre-shared-key",
				"pre-shared-key-xauth", "rsa-key", "rsa-signature-hybrid"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of a certificate listed in System/Certificates (signing packets; the certificate must " +
				"have the private key). Applicable if digital signature authentication method (`auth-method=digital-signature`) " +
				"or EAP (a`uth-method=eap`) is used.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"eap_methods": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "All EAP methods requires whole certificate chain including intermediate and root CA certificates " +
				"to be present in System/Certificates menu. Also, the username and password (if required by the authentication " +
				"server) must be specified. Multiple EAP methods may be specified and will be used in a specified order. " +
				"Currently supported EAP methods:\n  * eap-mschapv2;\n  * eap-peap - also known as PEAPv0/EAP-MSCHAPv2;\n  * eap-tls - " +
				"requires additional client certificate specified under certificate parameter;\n  * eap-ttls.",
			ValidateFunc:     validation.StringInSlice([]string{"eap-mschapv2", "eap-peap", "eap-tls", "eap-ttls"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"generate_policy": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allow this peer to establish SA for non-existing policies. Such policies are created dynamically " +
				"for the lifetime of SA. Automatic policies allows, for example, to create IPsec secured L2TP tunnels, " +
				"or any other setup where remote peer's IP address is not known at the configuration time. `no` - do not " +
				"generate policies; `port-override` - generate policies and force policy to use any port (old behavior); " +
				"`port-strict` - use ports from peer's proposal, which should match peer's policy.",
			ValidateFunc:     validation.StringInSlice([]string{"no", "port-override", "port-strict"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the private key from keys menu. Applicable if RSA key authentication method (`auth-method=rsa-key`) " +
				"is used.",
		},
		"match_by": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Defines the logic used for peer's identity validation. `remote-id` - will verify the peer's ID " +
				"according to remote-id setting. `certificate` will verify the peer's certificate with what is specified " +
				"under remote-certificate setting.",
			ValidateFunc:     validation.StringInSlice([]string{"remote-id", "certificate"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mode_config": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the configuration parameters from mode-config menu. When parameter is set mode-config " +
				"is enabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"my_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "On initiator, this controls what ID_i is sent to the responder. On responder, this controls " +
				"what ID_r is sent to the initiator. In IKEv2, responder also expects this ID in received ID_r from initiator. `auto` " +
				"- tries to use correct ID automatically: IP for pre-shared key, SAN (DN if not present) for certificate " +
				"based connections; `address` - IP address is used as ID;dn - the binary Distinguished Encoding Rules (DER) " +
				"encoding of an ASN.1 X.500 Distinguished Name; `fqdn` - fully qualified domain name; `key-id` - use the specified " +
				"key ID for the identity; `user-fqdn` - specifies a fully-qualified username string, for example, `user@domain.com`.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "address", "fqdn", "user-fqdn", "key-id"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"notrack_chain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Adds IP/Firewall/Raw rules matching IPsec policy to a specified chain. Use together with generate-policy.",
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "XAuth or EAP password. Applicable if pre-shared key with XAuth authentication method " +
				"(`auth-method=pre-shared-key-xauth`) or EAP (`auth-method=eap`) is used.",
		},
		"peer": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the peer on which the identity applies.",
		},
		"policy_template_group": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If generate-policy is enabled, traffic selectors are checked against templates from the same " +
				"group. If none of the templates match, Phase 2 SA will not be established.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of a certificate (listed in `System/Certificates`) for authenticating the remote side (validating " +
				"packets; no private key required). If a remote-certificate is not specified then the received certificate " +
				"from a remote peer is used and checked against CA in the certificate menu. Proper CA must be imported " +
				"in a certificate store. If remote-certificate and match-by=certificate is specified, only the specific " +
				"client certificate will be matched. Applicable if digital signature authentication method " +
				"(`auth-method=digital-signature`) is used.",
		},
		"remote_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This parameter controls what ID value to expect from the remote peer. Note that all types " +
				"except for ignoring will verify remote peer's ID with a received certificate. In case when the peer " +
				"sends the certificate name as its ID, it is checked against the certificate, else the ID is checked " +
				"against Subject Alt. Name. `auto` - accept all ID's;address - IP address is used as ID;dn - the binary " +
				"Distinguished Encoding Rules (DER) encoding of an ASN.1 X.500 Distinguished Name; `fqdn` - fully qualified " +
				"domain name. Only supported in IKEv2; `user-fqdn` - a fully-qualified username string, for example, `user@domain.com`. " +
				"Only supported in IKEv2; `key-id` - specific key ID for the identity. Only supported in IKEv2; `ignore` - " +
				"do not verify received ID with certificate (dangerous). * Wildcard key ID matching **is not supported**, " +
				"for example `remote-id=`key-id:CN=*.domain.com`.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "fqdn", "user-fqdn", "key-id", "ignore"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the public key from keys menu. Applicable if RSA key authentication method " +
				"(`auth-method=rsa-key`) is used.",
		},
		"secret": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Secret string. If it starts with '0x', it is parsed as a hexadecimal value. Applicable if " +
				"pre-shared key authentication method (`auth-method=pre-shared-key` and `auth-method=pre-shared-key-xauth`) " +
				"is used.",
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "XAuth or EAP username. Applicable if pre-shared key with XAuth authentication method " +
				"(`auth-method=pre-shared-key-xauth`) or EAP (`auth-method=eap`) is used.",
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

		Schema: resSchema,
	}
}
