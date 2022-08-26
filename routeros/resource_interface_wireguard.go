package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*55",
    "disabled": "false",
    "listen-port": "13231",
    "mtu": "1420",
    "name": "wireguard1",
    "private-key": "gLP306E2BCZBeyZ0ILrS5Ubdg4VjkFYiWkg7HpKYM10=",
    "public-key": "HhbyDzG6loyFAsB040GvnOcRH1Ks+M44utp6REaWPxo=",
    "running": "true"
  }
*/

// ResourceInterfaceWireguard https://help.mikrotik.com/docs/display/ROS/WireGuard
func ResourceInterfaceWireguard() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireguard"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"listen_port": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "Port for WireGuard service to listen on for incoming sessions.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		KeyMtu:  PropMtuRw(1420),
		KeyName: PropNameRw,
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
