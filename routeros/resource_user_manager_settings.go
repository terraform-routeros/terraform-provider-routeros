package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    "accounting-port": "1813",
    "authentication-port": "1812",
    "certificate": "*0",
    "enabled": "true",
    "use-profiles": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Settings
func ResourceUserManagerSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager"),
		MetaId:           PropId(Id),

		"accounting_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1813,
			Description:  "Port to listen for RADIUS accounting requests.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"authentication_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1812,
			Description:  "Port to listen for RADIUS authentication requests.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Certificate for use in EAP TLS-type authentication methods.",
		},
		KeyEnabled: PropEnabled("An option whether the User Manager functionality is enabled."),
		"require_message_auth": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option whether to require `Message-Authenticator` in received Access-Accept/Challenge/Reject messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"no", "yes-access-request"}, false),
		},
		"use_profiles": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to use Profiles and Limitations. When set to `false`, only User configuration is required to run User Manager.",
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
