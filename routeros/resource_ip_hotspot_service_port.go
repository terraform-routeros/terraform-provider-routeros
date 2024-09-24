package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "name": "ftp",
    "ports": "21"
  }
*/

func ResourceIpHotspotServicePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/service-port"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("name"),

		KeyDisabled: PropDisabledRw,
		KeyName:     PropName("Service name."),
		"ports": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreateUpdate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultCreateUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
