package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*4",
    "action": "echo",
    "default": "true",
    "disabled": "false",
    "invalid": "false",
    "prefix": "",
    "topics": "critical"
}
*/

func ResourceSystemLogging() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/logging"),
		MetaId:           PropId(Id),
		"action": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "specifies one of the system default actions or user specified action listed in actions menu",
			ValidateFunc: validation.StringInSlice([]string{"disk", "echo", "memory", "remote"}, false),
		},
		"prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "prefix added at the beginning of log messages",
			Default:     "",
		},
		KeyDisabled: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether or not this logging should be disabled",
		},
		"invalid": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"topics": {
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "prefix added at the beginning of log messages",
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
