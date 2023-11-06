package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceInterfaceWireguardV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/wireguard"),
			MetaId:           PropId(Name),

			KeyComment:  PropCommentRw,
			KeyDisabled: PropDisabledRw,
			"listen_port": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Port for WireGuard service to listen on for incoming sessions.",
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			KeyMtu:  PropMtuRw(),
			KeyName: PropNameForceNewRw,
			"private_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
				Description: "A base64 private key. If not specified, it will be automatically " +
					"generated upon interface creation.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A base64 public key is calculated from the private key.",
			},
			KeyRunning: PropRunningRo,
		},
	}
}
