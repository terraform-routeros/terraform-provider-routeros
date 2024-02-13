package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Addenvironmentvariablesandmounts(optional)
func ResourceContainerMounts() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container/mounts"),
		MetaId:           PropId(Name),

		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the mount.",
		},
		"src": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specifies source path of the mount, which points to a RouterOS location",
		},
		"dst": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specifies destination path of the mount, which points to defined location in container",
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
