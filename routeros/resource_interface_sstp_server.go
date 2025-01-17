package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "authentication": "pap,chap,mschap1,mschap2",
  "certificate": "none",
  "ciphers": "aes256-sha,aes256-gcm-sha384",
  "default-profile": "default",
  "enabled": "false",
  "keepalive-timeout": "60",
  "max-mru": "1500",
  "max-mtu": "1500",
  "mrru": "disabled",
  "pfs": "false",
  "port": "443",
  "tls-version": "any",
  "verify-client-certificate": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/SSTP#SSTP-SSTPServer
func ResourceInterfaceSSTPServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/sstp-server/server"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:             schema.TypeSet,
			Optional:         true,
			Description:      "Authentication algorithm.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				Default:      "all",
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
		},
		"keepalive_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "60",
			Description:      "Sets keepalive timeout in seconds.",
			DiffSuppressFunc: TimeEquall,
		},
		"port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets port used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the certificate in use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
		"tls_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which TLS versions to allow.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"any", "only-1.2"}, false),
		},
		"ciphers": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      `Allowed ciphers.`,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"null", "aes256-sha", "aes256-gcm-sha384"}, false, false),
		},
		"verify_client_certificate": {
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Description:      "SSTP server will verify client certificate.",
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, " +
				"it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the " +
				"tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Default profile to use.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyEnabled: PropEnabled("Enables/disables service."),
		"pfs": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Specifies which TLS authentication to use. With pfs=yes, TLS will use ECDHE-RSA- and DHE-RSA-. " +
				"For maximum security setting pfs=required will use only ECDHE.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
