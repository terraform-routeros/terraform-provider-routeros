package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    "db-path": "/flash/user-manager5",
    "db-size": "78176",
    "free-disk-space": "3248128"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Database
func ResourceUserManagerDatabase() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/database"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields(`"db_size","free_disk_space"`),

		"db_path": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Path to the location where database files will be stored.",
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
