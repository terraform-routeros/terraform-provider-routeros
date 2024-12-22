package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "allow-none-crypto": "false",
  "always-allow-password-login": "false",
  "forwarding-enabled": "no",
  "host-key-size": "2048",
  "strong-crypto": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/SSH#SSH-SSHServer
func ResourceIpSSHServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ssh"),
		MetaId:           PropId(Id),

		"allow_none_crypto": {
			Type:         schema.TypeBool,
			Optional:     true,
			Description:  "Whether to allow connection if cryptographic algorithms are set to none.",
			ExactlyOneOf: []string{"allow_none_crypto", "strong_crypto"},
		},
		"always_allow_password_login": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to allow password login at the same time when public key authorization is " +
				"configured for a user.",
		},
		"forwarding_enabled": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows to control which SSH forwarding method to allow:" +
				"\n  * no - SSH forwarding is disabled;" +
				"\n  * local - Allow SSH clients to originate connections from the server(router), this setting controls also dynamic forwarding;" +
				"\n  * remote - Allow SSH clients to listen on the server(router) and forward incoming connections;" +
				"\n  * both - Allow both local and remote forwarding methods.",
			ValidateFunc:     validation.StringInSlice([]string{"both", "local", "no", "remote"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"host_key_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "RSA key size when host key is being regenerated.",
			ValidateFunc:     validation.IntInSlice([]int{1024, 1536, 2048, 4096, 8192}),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"host_key_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Select host key type.",
			ValidateFunc:     validation.StringInSlice([]string{"rsa", "ed25519"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"strong_crypto": {
			Type:         schema.TypeBool,
			Optional:     true,
			Description:  "Use stronger encryption.",
			ExactlyOneOf: []string{"allow_none_crypto", "strong_crypto"},
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
