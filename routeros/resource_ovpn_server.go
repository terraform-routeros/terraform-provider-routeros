package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    "auth": "sha1,md5,sha256,sha512",
    "certificate": "root-cert",
    "cipher": "blowfish128,aes128-cbc",
    "default-profile": "default",
    "enable-tun-ipv6": "false",
    "enabled": "true",
    "ipv6-prefix-len": "64",
    "keepalive-timeout": "60",
    "mac-address": "FE:01:63:24:35:19",
    "max-mtu": "1500",
    "mode": "ip",
    "netmask": "24",
    "port": "1194",
    "protocol": "tcp",
    "redirect-gateway": "disabled",
    "reneg-sec": "3600",
    "require-client-certificate": "true",
    "tls-version": "only-1.2",
    "tun-server-ipv6": "::"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/OpenVPN
func ResourceOpenVPNServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ovpn-server/server"),
		MetaId:           PropId(Id),

		"auth": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "null", "sha256", "sha512"}, false),
			},
			Description:      "Authentication methods that the server will accept.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the certificate that the OVPN server will use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cipher": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{
					"null", "aes128-cbc", "aes128-gcm", "aes192-cbc", "aes192-gcm", "aes256-cbc", "aes256-gcm", "blowfish128",
					// Backward compatibility with ROS v7.7
					"aes128", "aes192", "aes256",
				}, false),
			},
			Description:      `Allowed ciphers.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			Description: "Default profile to use.",
		},
		"enable_tun_ipv6": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Specifies if IPv6 IP tunneling mode should be possible with this OVPN server.",
		},
		KeyEnabled: PropEnabled("Defines whether the OVPN server is enabled or not."),
		"ipv6_prefix_len": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
			Description: "Length of IPv6 prefix for IPv6 address which will be used when generating OVPN interface " +
				"on the server side.",
			ValidateFunc: validation.IntBetween(1, 128),
		},
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "60",
			Description: "Defines  the time period (in seconds) after which the router is starting to send  " +
				"keepalive packets every second. If no traffic and no keepalive  responses have come for " +
				"that period of time (i.e. 2 *  keepalive-timeout), not responding client is proclaimed " +
				"disconnected",
			DiffSuppressFunc: TimeEquall,
		},
		// Computed only???
		"mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Automatically generated MAC address of the server.",
		},
		"max_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
			Description: "Maximum Transmission Unit. Max packet size that the OVPN interface will be able to send " +
				"without packet fragmentation.",
			ValidateFunc: validation.IntBetween(64, 65535),
		},
		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "ip",
			Description:  "Layer3 or layer2 tunnel mode (alternatively tun, tap)",
			ValidateFunc: validation.StringInSlice([]string{"ip", "ethernet"}, false),
		},
		"netmask": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      24,
			Description:  "Subnet mask to be applied to the client.",
			ValidateFunc: validation.IntBetween(0, 32),
		},
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1194,
			Description:  "Port to run the server on.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "tcp",
			Description:  "indicates the protocol to use when connecting with the remote endpoint.",
			ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
		},
		"push_routes": {
			Type:             schema.TypeSet,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Push routes to the VPN client (available since RouterOS 7.14).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"redirect_gateway": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"def1", "disabled", "ipv6"}, false),
			},
			Description: "Specifies what kind of routes the OVPN client must add to the routing table.\n  * def1 – Use " +
				"this flag to override the default gateway by using 0.0.0.0/1 and  128.0.0.0/1 rather " +
				"than 0.0.0.0/0. This has the benefit of overriding  but not wiping out the original " +
				"default gateway.\n  * disabled - Do not send redirect-gateway flags to the OVPN client.\n  * ipv6 " +
				"- Redirect IPv6 routing into the tunnel on the client side. This works  similarly to the " +
				"def1 flag, that is, more specific IPv6 routes are added  (2000::/4 and 3000::/4), " +
				"covering the whole IPv6 unicast space.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"reneg_sec": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     3600,
			Description: "Renegotiate data channel key after n seconds (default=3600).",
		},
		"require_client_certificate": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "If set to yes, then the server checks whether the client's certificate belongs to the " +
				"same certificate chain.",
		},
		"tls_version": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "any",
			Description:  "Specifies which TLS versions to allow.",
			ValidateFunc: validation.StringInSlice([]string{"any", "only-1.2"}, false),
		},
		"tun_server_ipv6": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "::",
			Description: "IPv6 prefix address which will be used when generating the OVPN interface on the server " +
				"side.",
		},
	}

	return &schema.Resource{
		Description:   `##### *<span style="color:red">This resource requires a minimum version of RouterOS 7.8!</span>*`,
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema:        resSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceOpenVPNServerV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationScalarToList("auth", "cipher", "redirect_gateway"),
				Version: 0,
			},
		},
	}
}
