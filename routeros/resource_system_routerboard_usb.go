package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "type": "auto",
  "usb-mode": "automatic"
}
*/

// https://help.mikrotik.com/docs/display/ROS/USB+Features
func ResourceSystemRouterboardUsb() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/routerboard/usb"),
		MetaId:           PropId(Id),

		"type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set the type of the USB port. Possible value: `auto`, `mini-PCIe`, `USB-type-A`.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "mini-PCIe", "USB-type-A"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"usb_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set the USB port mode. Possible values: `automatic`, `force-host`.",
			ValidateFunc:     validation.StringInSlice([]string{"automatic", "force-host"}, false),
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
