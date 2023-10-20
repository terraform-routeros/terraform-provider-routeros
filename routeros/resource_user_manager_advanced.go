package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    "paypal-allow": "false",
    "paypal-currency": "USD",
    "paypal-password": "",
    "paypal-signature": "",
    "paypal-use-sandbox": "false",
    "paypal-user": "",
    "web-private-password": "",
    "web-private-username": ""
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Advanced
func ResourceUserManagerAdvanced() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/advanced"),
		MetaId:           PropId(Id),

		"paypal_allow": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to enable PayPal functionality for User Manager.",
		},
		"paypal_currency": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "USD",
			Description: "The currency related to price setting in which users will be billed.",
		},
		"paypal_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The password of the PayPal API account.",
		},
		"paypal_signature": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The signature of the PayPal API account.",
		},
		"paypal_use_sandbox": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to use PayPal's sandbox environment for testing purposes.",
		},
		"paypal_user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username of the PayPal API account.",
		},
		"web_private_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The password for accessing `/um/PRIVATE/` section over HTTP.",
		},
		"web_private_username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username for accessing `/um/PRIVATE/` section over HTTP.",
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
