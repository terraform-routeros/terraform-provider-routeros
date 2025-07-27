package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
	"auto-smb-sharing": "false",
	"auto-smb-user": "guest",
	"auto-media-sharing": "false",
	"auto-media-interface": "lo",
	"default-mount-point-template": "[slot]"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/91193346/Disks#Disks-Settings
func ResourceDiskSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/disk/settings"),
		MetaId:           PropId(Id),

		"auto_smb_sharing": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables dynamic SMB shares when new disk/partition item is added in '/disk'.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"auto_smb_user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Default value for smb-sharing/smb-user setting, when new disk/partition item is added in '/disk'.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"auto_media_sharing": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables media dynamically when new disk/partition item is added in '/disk'.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"auto_media_interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interface that will be used in dynamic instance for ip/media when new disk/partition item is added in '/disk'.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_mount_point_template": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Sets the default mount point template for each item added in `/disk`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
