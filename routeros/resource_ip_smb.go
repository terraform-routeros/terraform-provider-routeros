package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
	"enabled": "auto",
	"status": "disabled",
	"domain": "MSHOME",
	"comment": "MikrotikSMB",
	"interfaces": "all"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/117145608/SMB
func ResourceIpSMB() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/smb"),
		MetaId:           PropId(Id),

		KeyEnabled: {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The default value is 'auto'. This means that the SMB server will automatically be enabled when the first non-disabled SMB share is configured under '/ip smb share'.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     ValidationAutoYesNo,
		},
		"domain": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of Windows Workgroup.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: {
			// This is the SMB server comment, not a MikroTik comment, which is why this deserves its own attribute and not PropCommentRw.
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Set comment for the server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interfaces": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:         true,
			Description:      "List of interfaces on which SMB service will be running.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
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
