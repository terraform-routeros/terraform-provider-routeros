package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "address": "127.0.0.1",
    "coa-port": "3799",
    "disabled": "false",
    "name": "test",
    "shared-secret": "password"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Routers
func ResourceUserManagerRouter() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/router"),
		MetaId:           PropId(Id),

		"address": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "IP address of the RADIUS client.",
			ValidateFunc: ValidationIpAddress,
		},
		"coa_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      3799,
			Description:  "Port number of CoA (Change of Authorization) communication.",
			ValidateFunc: Validation64k,
		},
		KeyDisabled: PropDisabledRw,
		KeyName:     PropName("Unique name of the RADIUS client."),
		"shared_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The shared secret to secure communication.",
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
