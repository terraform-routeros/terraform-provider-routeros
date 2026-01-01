package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Addenvironmentvariablesandmounts(optional)
func ResourceContainerMounts() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container/mounts"),
		MetaId:           PropId(Id),

		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID for the mount",
		},
		"disabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the mount is disabled",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"list": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Mount list this mount is for",
		},
		"read_only": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to mount read-only",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
