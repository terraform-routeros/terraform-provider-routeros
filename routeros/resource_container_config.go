package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Containerconfiguration
func ResourceContainerConfig() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container/config"),
		MetaId:           PropId(Name),

		"registry_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "External registry url from where the container will be downloaded.",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the username for authentication (starting from ROS 7.8)",
		},
		"password": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Optional:    true,
			Description: "Specifies the password for authentication (starting from ROS 7.8)",
		},
		"ram_high": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0",
			Description: "RAM usage limit. (0 for unlimited)",
		},
		"tmpdir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container extraction directory.",
		},
		"layer_dir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container layers directory.",
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
