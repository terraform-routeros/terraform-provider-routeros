package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServerConfig https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-StoreConfiguration
func ResourceDhcpServerConfig() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/config"),
		MetaId:           PropId(Id),

		"accounting": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "An option that enables accounting for DHCP leases.",
		},
		"interim_update": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "0s",
			Description:      "An option determining whether the DHCP server sends periodic updates to the accounting server during a lease.",
			DiffSuppressFunc: TimeEqual,
		},
		"radius_password": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "empty",
			Description:  "An option to set the password parameter for the RADIUS server. This option is available in RouterOS starting from version 7.0.",
			ValidateFunc: validation.StringInSlice([]string{"empty", "same-as-user"}, false),
		},
		"store_leases_disk": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "5m",
			Description:      "An option of how often the DHCP leases will be stored on disk.",
			DiffSuppressFunc: TimeEqual,
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
