package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
	".id": "*1",
	"add-default-route": "false",
	"add-sni": "false",
	"authentication": "pap,chap,mschap1,mschap2",
	"certificate": "none",
	"ciphers": "aes256-sha",
	"connect-to": "sstp.example.com",
	"dial-on-demand": "false",
	"disabled": "false",
	"http-proxy": "::",
	"hw-crypto": "false",
	"keepalive-timeout": "60",
	"max-mru": "1500",
	"max-mtu": "1500",
	"mrru": "disabled",
	"name": "sstp-out1",
	"password": "password",
	"pfs": "false",
	"port": "443",
	"profile": "default-encryption",
	"proxy-port": "443",
	"running": "false",
	"tls-version": "any",
	"user": "username",
	"verify-server-address-from-certificate": "true",
	"verify-server-certificate": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/SSTP#SSTP-SSTPClient
func ResourceInterfaceSSTPClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/sstp-client"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Authentication algorithm.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				Default:      "all",
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"add_default_route": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to add L2TP remote address as a default route.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_route_distance": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets distance value applied to auto created default route, if add-default-route is also selected.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			RequiredWith:     []string{"add_default_route"},
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, " +
				"it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the " +
				"tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"proxy_port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets proxy port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"add_sni": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables/disables service.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dial_on_demand": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Connects only when outbound traffic is generated. If selected, then route with gateway address " +
				"from 10.112.112.0/24 network will be added while connection is not established.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName:    PropName("Descriptive name of the interface."),
		KeyRunning: PropRunningRo,
		"hw_crypto": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},
		"tls_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which TLS versions to allow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"any", "only-1.2"}, false),
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "User name used for authentication.",
		},
		"certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the certificate in use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"http_proxy": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Proxy address field.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password used for authentication.",
			Sensitive:   true,
		},
		"verify_server_address_from_certificate": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "SSTP client will verify server address in certificate.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"verify_server_certificate": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "SSTP client will verify server certificate.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ciphers": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Allowed ciphers.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"null", "aes256-sha", "aes256-gcm-sha384"}, false, false),
		},
		"keepalive_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "60",
			Description:      "Sets keepalive timeout in seconds.",
			DiffSuppressFunc: TimeEquall,
		},
		"pfs": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies which TLS authentication to use. With pfs=yes, TLS will use ECDHE-RSA- and DHE-RSA-. " +
				"For maximum security setting pfs=required will use only ECDHE.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"max_mru": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Receive Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(512, 18432),
		},
		"max_mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Transmission Unit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntBetween(512, 18432),
		},
		"port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets port used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"connect_to": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Remote address of the SSTP server.",
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which PPP profile configuration will be used when establishing the tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
