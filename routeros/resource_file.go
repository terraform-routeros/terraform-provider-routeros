package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/pages/viewpage.action?pageId=2555971
func ResourceFile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/file"),
		MetaId:           PropId(Id),

		"contents": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "The actual content of the file",
		},
		"creation_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A time when the file was created",
		},
		"last_modified": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A time when the file was modified",
		},
		KeyName: PropName("Name of the file"),
		"package_architecture": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Architecture that package is built for. Applies only to RouterOS \".npk\" files",
		},
		"package_built_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A time when the package was built. Applies only to RouterOS \".npk\" files",
		},
		"package_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the installable package. Applies only to RouterOS \".npk\" files",
		},
		"package_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A version of the installable package. Applies only to RouterOS \".npk\" files",
		},
		"size": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "File size in bytes",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Type of the file. For folders, the file type is the directory",
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
