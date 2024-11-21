package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "current-firmware":"7.11.2",
  "factory-firmware":"6.48.6",
  "firmware-type":"qca9531L",
  "model":"CRS312-4C+8XG",
  "revision":"r2",
  "routerboard":"true",
  "serial-number":"XXXXXXXXXX",
  "upgrade-firmware":"7.11.2"
}
*/

// https://help.mikrotik.com/docs/display/ROS/
func DatasourceSystemRouterboard() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/routerboard"),
		MetaId:           PropId(Id),

		"current_firmware": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"factory_firmware": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"firmware_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"model": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"revision": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"routerboard": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"upgrade_firmware": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		ReadContext: DefaultSystemDatasourceRead(resSchema),
		Schema:      resSchema,
	}
}
